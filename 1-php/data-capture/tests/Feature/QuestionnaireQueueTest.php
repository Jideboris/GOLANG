<?php

namespace Tests\Feature;

use App\Jobs\QuestionnaireCompleted;
use App\Models\Questionnaire;
use App\Models\QuestionnaireResult;
use App\Models\ScheduledQuestionnaire;
use App\Models\User;
use Tests\TestCase;

class QuestionnaireQueueTest extends TestCase
{
    /** @test */
    public function it_can_push_message_to_queue()
    {
        // given
        $participant = User::factory()->create();
        $questionnaire = Questionnaire::factory()->create();
        $schedule = ScheduledQuestionnaire::factory()->create([
            'questionnaire_id'          => $questionnaire->id,
            'participant_id'            => $participant->id
        ]);

        $result = QuestionnaireResult::factory()->create([
            'questionnaire_id'          => $questionnaire->id,
            'participant_id'            => $participant->id,
            'questionnaire_schedule_id' => $schedule->id
        ]);

        // when
        $job = new QuestionnaireCompleted($result);

        // then
        $this->assertObjectHasAttribute('id', $job);
        
        $this->assertObjectHasAttribute('userId', $job);
        
        $this->assertObjectHasAttribute('studyId', $job);
        
        $this->assertObjectHasAttribute('questionnaireId', $job);
        
        $this->assertObjectHasAttribute('completedAt', $job);
    }
}
