<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateQuestionnaireResultsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('questionnaire_results', function (Blueprint $table) {
            $table->uuid('id')->primary();

            $table->json('answers');
            $table->uuid('questionnaire_id');
            $table->uuid('participant_id');
            $table->uuid('questionnaire_schedule_id')->nullable();

            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('questionnaire_results');
    }
}
