<?php

namespace App\Http\Controllers;

use App\Models\Questionnaire;
use App\Models\QuestionnaireResult;
use App\Models\ScheduledQuestionnaire;
use App\Services\QuestionnaireResultService;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;

class QuestionnaireController extends Controller
{
    /**
     * Handles the submission of a new questionnaire result.
     * 
     * @param Request $request
     */
    public function submitResult(Request $request)
    {
        // validate request & return messages as json
        $validator = Validator::make($request->all(), [
            'questionnaire_id'          => 'required|exists:questionnaires,id',
            'results'                   => 'required',
            'questionnaire_schedule_id' => 'exists:scheduled_questionnaires,id'
        ]);

        if ($validator->fails()) {
            return response()->json($validator->messages(), 400);
        }

        // get current authed user
        $user = auth()->user();

        // create new result
        $result = QuestionnaireResultService::createResult(
            $request->input('questionnaire_id'),
            $user->id,
            json_encode($request->input('results')),
            $request->input('questionnaire_schedule_id')
        );

        // return OK
        return response()->json([
            'status' => 200,
            'result' => $result
        ]);
    }
}
