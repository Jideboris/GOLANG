<?php

namespace Tests\Feature;

use App\Models\Questionnaire;
use App\Models\QuestionnaireResult;
use App\Models\ScheduledQuestionnaire;
use App\Models\User;
use App\Services\QuestionnaireResultService;
use App\Services\QueueService;
use Tests\TestCase;
use Aws\Sqs\SqsClient;
use Mockery\MockInterface;

class QuestionnaireQueueTest extends TestCase
{
    /** @test */
    public function it_can_push_message_to_queue()
    {
        // given
        $participant = User::factory()->create();
        $questionnaire = Questionnaire::factory()->create();
        $schedule = ScheduledQuestionnaire::factory()->create([
            'questionnaire_id'          => $questionnaire->id,
            'participant_id'            => $participant->id
        ]);

        $result = QuestionnaireResult::factory()->create([
            'questionnaire_id'          => $questionnaire->id,
            'participant_id'            => $participant->id,
            'questionnaire_schedule_id' => $schedule->id
        ]);

        $clientMock = $this->getMockBuilder(SqsClient::class)
            ->disableOriginalConstructor()
            ->addMethods([ 'sendMessage' ])
            ->getMock();

        $clientMock->expects($this->any())
            ->method('sendMessage')
            ->willReturn([
                '@metadata' => [
                    'statusCode' => 200
                ]
            ]);

        $service = $this->partialMock(QueueService::class, function (MockInterface $mock) use ($clientMock) {
            $mock->shouldReceive('getClient')
                ->once()
                ->andReturn($clientMock);
        });

        // when
        $result = $service->pushToQueue($result);

        // then
        $this->assertTrue($result);
    }
}
