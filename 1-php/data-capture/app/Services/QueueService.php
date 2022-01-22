<?php

namespace App\Services;

use App\Models\QuestionnaireResult;
use Aws\Exception\AwsException;
use Aws\Sqs\SqsClient;

class QueueService
{
    /**
     * Gets the SQS client.
     */
    public function getClient()
    {
        return new SqsClient([
            'region'        => 'eu-west-1',
            'version'       => 'latest',
            'credentials'   => [
                'key'    => env('AWS_ACCESS_KEY_ID', ''),
                'secret' => env('AWS_SECRET_ACCESS_KEY', ''),
            ]
        ]);
    }

    /**
     * Pushes questionnaire result into sqs queue.
     * 
     * @param QuestionnaireResult $result
     * @return bool
     * @throws AwsException
     */
    public function pushToQueue(QuestionnaireResult $result): bool
    {
        $client = $this->getClient();

        $params = [
            'DelaySeconds' => 10,
            'MessageAttributes' => [
                'Id' => [
                    'DataType' => 'String',
                    'StringValue' => $result->id
                ],
                'UserId' => [
                    'DataType' => 'String',
                    'StringValue' => $result->participant_id
                ],
                'StudyId' => [
                    'DataType' => 'String',
                    'StringValue' => $result->questionnaire->study_id
                ],
                'QuestionnaireId' => [
                    'DataType' => 'String',
                    'StringValue' => $result->questionnaire->id
                ],
                'CompletedAt' => [
                    'DataType' => 'String',
                    'StringValue' => $result->created_at->toDateTimeString()
                ]
            ],
            'MessageBody' => "Information about the questionnaire result.",
            'QueueUrl' => env('SCHEDULE_QUEUE_URL')
        ];

        $result = $client->sendMessage($params);
        
        return isset($result['@metadata']['statusCode']) && $result['@metadata']['statusCode'] === 200;
    }
}