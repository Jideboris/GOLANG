<?php

namespace App\Jobs;

use App\Models\QuestionnaireResult;
use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Queue\SerializesModels;

class QuestionnaireCompleted implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    protected $id;

    protected $userId;

    protected $studyId;

    protected $questionnaireId;

    protected $completedAt;

    /**
     * Create a new job instance.
     *
     * @return void
     */
    public function __construct(QuestionnaireResult $result)
    {
        $this->id = $result->id;
        $this->userId = $result->participant_id;
        $this->studyId = $result->questionnaire->study_id;
        $this->questionnaireId = $result->questionnaire_id;
        $this->completedAt = $result->created_at->toDateTimeString();
    }

    /**
     * Execute the job.
     *
     * @return void
     */
    public function handle()
    {
        
    }
}
