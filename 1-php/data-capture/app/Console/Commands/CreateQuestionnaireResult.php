<?php

namespace App\Console\Commands;

use App\Models\Questionnaire;
use App\Models\QuestionnaireResult;
use App\Models\ScheduledQuestionnaire;
use App\Models\User;
use App\Services\QuestionnaireResultService;
use Illuminate\Console\Command;
use Illuminate\Contracts\Filesystem\FileNotFoundException;

class CreateQuestionnaireResult extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'questionnaire_result:create {user_id} {questionnaire_id} {json_file} {--schedule=}';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Creates a new questionnaire result.';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return int
     */
    public function handle()
    {
        $userArgument = $this->argument('user_id');
        $questionnaireArgument = $this->argument('questionnaire_id');
        $jsonFileArgument = storage_path($this->argument('json_file'));
        $scheduleOption = $this->option('schedule');

        if (! file_exists($jsonFileArgument)) {
            throw new FileNotFoundException();
        }

        $user = User::findOrFail($userArgument);
        $questionnaire = Questionnaire::findOrFail($questionnaireArgument);

        QuestionnaireResultService::createResult(
            $questionnaire->id,
            $user->id,
            file_get_contents($jsonFileArgument),
            $scheduleOption
        );

        $this->info('Result successfully created.');
    }
}
