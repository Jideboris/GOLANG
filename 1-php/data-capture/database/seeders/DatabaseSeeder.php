<?php

namespace Database\Seeders;

use App\Models\User;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     *
     * @return void
     */
    public function run()
    {
        $user = User::factory()->create();

        DB::table('scheduled_questionnaires')->insert([
            [
                'id'                => '02f86fb2-763a-4001-b7a0-3f8714e0099d',
                'questionnaire_id'  => 'e147be89-0116-4aec-8d38-b9af4d51817b',
                'participant_id'    => $user->id,
                'scheduled_at'      => '2021-01-03 12:00:00',
                'status'            => 'complete',
                'created_at'        => now(),
                'updated_at'        => now()
            ]
        ]);
        
        DB::table('questionnaires')->insert([
            [
                'id'                        => 'e078de03-b277-4ce0-affd-20297b150e97',
                'study_id'                  => 'limbs-study',
                'name'                      => 'Who has legs?',
                'questions'                 => '{}',
                'max_attempts'              => 1,
                'hours_between_attempts'    => 0,
                'created_at'                => now(),
                'updated_at'                => now()
            ]
        ]);
    }
}
