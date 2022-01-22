<?php

namespace Tests\Feature;

use App\Jobs\QuestionnaireCompleted;
use App\Models\Questionnaire;
use App\Models\ScheduledQuestionnaire;
use App\Models\User;
use Illuminate\Contracts\Filesystem\FileNotFoundException;
use Illuminate\Foundation\Testing\DatabaseTransactions;
use Illuminate\Support\Facades\Bus;
use Tests\TestCase;

class QuestionnaireCommandTest extends TestCase
{
    use DatabaseTransactions;

    protected $user;
    
    protected $questionnaire;

    protected $jsonPath = "/app/answers.json";

    public function setUp(): void
    {
        parent::setUp();

        $this->user = User::factory()->create();
        $this->questionnaire = Questionnaire::factory()->create();
    }
    
    /** @test */
    public function it_inserts_the_correct_data_into_the_database()
    {
        $this->artisan("questionnaire_result:create {$this->user->id} {$this->questionnaire->id} {$this->jsonPath}")
            ->assertSuccessful();

        $this->assertDatabaseHas('questionnaire_results', [
            'questionnaire_id'  => $this->questionnaire->id,
            'participant_id'    => $this->user->id
        ]);
    }
    
    /** @test */
    public function it_throws_an_exception_if_json_file_does_not_exist()
    {
        $this->expectException(FileNotFoundException::class);

        $this->artisan("questionnaire_result:create {$this->user->id} {$this->questionnaire->id} /app/noexistentfile.json")
            ->assertExitCode(0);
    }

    /** @test */
    public function it_accepts_schedule_id_parameter()
    {
        // given
        $schedule = ScheduledQuestionnaire::factory()->create([
            'questionnaire_id'  => $this->questionnaire->id,
            'participant_id'    => $this->user->id
        ]);

        Bus::fake();

        // when
        $this->artisan("questionnaire_result:create {$this->user->id} {$this->questionnaire->id} {$this->jsonPath} --schedule={$schedule->id}")
            ->assertSuccessful();

        // then
        $this->assertDatabaseHas('questionnaire_results', [
            'questionnaire_id'          => $this->questionnaire->id,
            'participant_id'            => $this->user->id,
            'questionnaire_schedule_id' => $schedule->id
        ]);

        $this->assertDatabaseHas('scheduled_questionnaires', [
            'id'        => $schedule->id,
            'status'    => 'complete'
        ]);

        Bus::assertDispatched(QuestionnaireCompleted::class);
    }
}
