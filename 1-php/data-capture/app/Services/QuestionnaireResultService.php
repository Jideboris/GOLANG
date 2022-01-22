<?php

namespace App\Services;

use App\Models\QuestionnaireResult;
use App\Models\ScheduledQuestionnaire;

class QuestionnaireResultService
{
    /**
     * Creates a new questionnaire result.
     * 
     * @param string $questionnaireID
     * @param string $participantID
     * @param string $answers
     * @param string|null $questionnaireScheduleID
     * @return QuestionnaireResult
     */
    public static function createResult(
        string $questionnaireID,
        string $participantID,
        string $answers,
        ?string $questionnaireScheduleID = null
    ): QuestionnaireResult
    {
        $result = new QuestionnaireResult();
        $result->questionnaire_id = $questionnaireID;
        $result->participant_id = $participantID;
        $result->answers = json_encode($answers);

        if (! is_null($questionnaireScheduleID)) {
            $scheduledQuestionnaire = ScheduledQuestionnaire::findOrFail($questionnaireScheduleID);
            $scheduledQuestionnaire->status = 'complete';
            $scheduledQuestionnaire->save();

            $result->questionnaire_schedule_id = $scheduledQuestionnaire->id;
        }

        $result->save();
        
        return $result;
    }
}