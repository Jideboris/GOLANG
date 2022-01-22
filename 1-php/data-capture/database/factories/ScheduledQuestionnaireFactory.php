<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

class ScheduledQuestionnaireFactory extends Factory
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
            'scheduled_at'              => now(),
            'status'                    => 'pending',
            'created_at'                => now(),
            'updated_at'                => now()
        ];
    }
}
