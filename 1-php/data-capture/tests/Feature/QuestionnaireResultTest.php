<?php

namespace Tests\Feature;

use App\Jobs\QuestionnaireCompleted;
use App\Models\Questionnaire;
use App\Models\ScheduledQuestionnaire;
use App\Models\User;
use App\Services\QuestionnaireResultService;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use Illuminate\Support\Facades\Bus;
use Illuminate\Testing\Fluent\AssertableJson;
use Symfony\Component\HttpKernel\Exception\MethodNotAllowedHttpException;
use Tests\TestCase;

class QuestionnaireResultTest extends TestCase
{
    use DatabaseTransactions;

    protected $user;

    public function setUp(): void
    {
        parent::setUp();

        $this->user = User::factory()->create();
    }

    /** @test */
    public function it_requires_authentication()
    {
        $response = $this->postJson('api/questionnaire_result');

        $response->assertStatus(401);
    }
        
    /** @test */
    public function it_requires_a_post_request()
    {
        $this->expectException(MethodNotAllowedHttpException::class);

        $this->withBasicAuth($this->user)
            ->withoutExceptionHandling()
            ->get('api/questionnaire_result');        
    }
    
    /** @test */
    public function it_validates_request()
    {
        $response = $this->withBasicAuth($this->user)
            ->postJson('api/questionnaire_result');

        $response->assertJson(function (AssertableJson $json) {
            $json->has('questionnaire_id')
                ->has('results');
        });
    }
    
    /** @test */
    public function it_inserts_the_correct_data_into_the_database()
    {
        // given
        $questionnaire = Questionnaire::factory()->create();

        // when
        $response = $this->withBasicAuth($this->user)
            ->postJson('api/questionnaire_result', [
                'questionnaire_id' => $questionnaire->id,
                'results' => [
                    'q1' => 'Hello World'
                ]
            ]);

        // then
        $this->assertDatabaseHas('questionnaire_results', [
            'questionnaire_id'  => $questionnaire->id,
            'participant_id'    => $this->user->id
        ]);
    }
    
    /** @test */
    public function it_validates_additional_questionnaire_schedule_id_parameter()
    {
        // given
        $questionnaire = Questionnaire::factory()->create();

        // when
        $response = $this->withBasicAuth($this->user)
            ->postJson('api/questionnaire_result', [
                'questionnaire_id' => $questionnaire->id,
                'questionnaire_schedule_id' => 'not_id',
                'results' => [
                    'q1' => 'Hello World'
                ]
            ]);

        // then
        $response->assertJson(function (AssertableJson $json) {
            $json->has('questionnaire_schedule_id');
        });
    }
    
    /** @test */
    public function it_links_result_to_schedule_and_updates_schedule_status()
    {
        // given
        $questionnaire = Questionnaire::factory()->create();
        $scheduledQuestionnaire = ScheduledQuestionnaire::factory()->create([
            'questionnaire_id'  => $questionnaire->id,
            'participant_id'    => $this->user->id
        ]);

        // when
        $response = $this->withBasicAuth($this->user)
            ->postJson('api/questionnaire_result', [
                'questionnaire_id' => $questionnaire->id,
                'questionnaire_schedule_id' => $scheduledQuestionnaire->id,
                'results' => [
                    'q1' => 'Hello World'
                ]
            ]);

        // then
        $this->assertDatabaseHas('questionnaire_results', [
            'questionnaire_id'          => $questionnaire->id,
            'participant_id'            => $this->user->id,
            'questionnaire_schedule_id' => $scheduledQuestionnaire->id
        ]);
        
        $this->assertDatabaseHas('scheduled_questionnaires', [
            'id'        => $scheduledQuestionnaire->id,
            'status'    => 'complete'
        ]);
    }

    /** @test */
    public function it_does_not_push_message_to_queue_when_result_has_no_schedule()
    {
        $participant = User::factory()->create();
        $questionnaire = Questionnaire::factory()->create();

        Bus::fake();

        QuestionnaireResultService::createResult(
            $questionnaire->id,
            $participant->id,
            '{ "q1": "Hello World" }'
        );

        Bus::assertNotDispatched(QuestionnaireCompleted::class);
    }
}
