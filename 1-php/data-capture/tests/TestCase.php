<?php

namespace Tests;

use App\Models\User;
use Aws\Sqs\SqsClient;
use Illuminate\Foundation\Testing\TestCase as BaseTestCase;

abstract class TestCase extends BaseTestCase
{
    use CreatesApplication;

    protected function withBasicAuth(User $user, $password = 'example'): self
    {
        return $this->withHeaders([
            'Authorization' => 'Basic ' . base64_encode("{$user->email}:{$password}")
        ]);
    }

    protected function mockSQSClient()
    {
        $clientMock = $this->getMockBuilder(SqsClient::class)
            ->disableOriginalConstructor()
            ->addMethods(['sendMessage'])
            ->getMock();

        $clientMock->expects($this->any())
            ->method('sendMessage')
            ->willReturn([
                '@metadata' => [
                    'statusCode' => 200
                ]
            ]);
        
        return $clientMock;
    }
}
