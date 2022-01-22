<?php

namespace Tests\Feature;

use App\Models\User;
use Exception;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Foundation\Testing\WithFaker;
use Symfony\Component\HttpKernel\Exception\MethodNotAllowedHttpException;
use Tests\TestCase;

class QuestionnaireResultTest extends TestCase
{
    protected $user;

    public function setUp(): void
    {
        parent::setUp();

        $this->user = User::factory()->create();
    }

    /** @test */
    public function it_requires_authentication()
    {
        $response = $this->post('api/questionnaire_result');

        $response->assertStatus(401);
    }
        
    /** @test */
    public function it_requires_a_post_request()
    {
        $this->expectException(MethodNotAllowedHttpException::class);

        $this->withBasicAuth($this->user)
            ->withoutExceptionHandling()
            ->get('api/questionnaire_result');        
    }
    
    /** @test */
    public function it_validates_request()
    {
        $response = $this->withBasicAuth($this->user)
            ->post('api/questionnaire_result');

        $response->assertInvalid([
            'questionnaire_id',
            'results'
        ]);
    }
    
    /** @test */
    public function it_inserts_the_correct_data_into_the_database()
    {
        
    }
    
    /** @test */
    public function it_takes_additional_questionnaire_schedule_id_parameter()
    {
        
    }
    
    /** @test */
    public function it_links_result_to_schedule_and_updates_schedule_status()
    {
        
    }
}
