<?php

namespace App\Http\Controllers;

use App\Models\Questionnaire;
use App\Models\QuestionnaireResult;
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
            'questionnaire_id' => 'required|exists:questionnaires,id',
            'results' => 'required'
        ]);

        if ($validator->fails()) {
            return response()->json($validator->messages(), 400);
        }

        // get current authed user
        $user = auth()->user();

        // create new questionnaire result
        $result = new QuestionnaireResult();
        $result->questionnaire_id = $request->input('questionnaire_id');
        $result->participant_id = $user->id;
        $result->questionnaire_schedule_id = '';
        $result->answers = json_encode($request->input('results'));
        $result->save();

        // return OK
        return response()->json([
            'status' => 200,
            'result' => $result
        ]);
    }
}
