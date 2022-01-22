<?php

namespace App\Http\Controllers;

use App\Http\Requests\QuestionnaireResultValidation;
use Illuminate\Http\Request;

class QuestionnaireController extends Controller
{
    /**
     * Handles the submission of a new questionnaire result.
     * 
     * @param Request $request
     */
    public function submitResult(QuestionnaireResultValidation $request)
    {
        return response()->json([
            'status' => 200
        ]);
    }
}
