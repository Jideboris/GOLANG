<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Str;

class QuestionnaireFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            'id'                        => Str::uuid(),
            'study_id'                  => 'limbs-study',
            'name'                      => 'Who has legs?',
            'questions'                 => '{}',
            'max_attempts'              => 1,
            'hours_between_attempts'    => 0,
            'created_at'                => now(),
            'updated_at'                => now()
        ];
    }
}
