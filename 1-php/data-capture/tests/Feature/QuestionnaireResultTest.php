<?php

namespace Tests\Feature;

use App\Models\Questionnaire;
use App\Models\User;
use Exception;
use Illuminate\Foundation\Testing\DatabaseTransactions;
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

        $response->assertJson(function (AssertableJson $json) {
            $json->has('questionnaire_schedule_id');
        });
    }
    
    /** @test */
    public function it_links_result_to_schedule_and_updates_schedule_status()
    {
        
    }
}
