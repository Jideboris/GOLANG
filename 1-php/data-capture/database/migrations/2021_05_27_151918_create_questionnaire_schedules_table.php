<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateQuestionnaireSchedulesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('scheduled_questionnaires', function (Blueprint $table) {
            $table->uuid('id')->primary();

            $table->uuid('questionnaire_id');
            $table->uuid('participant_id');
            $table->datetime('scheduled_at');
            $table->enum('status', ['pending', 'complete']);

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
        Schema::dropIfExists('scheduled_questionnaires');
    }
}
