package main

import (
    "log"
    "database/sql"
    "github.com/aws/aws-lambda-go/lambda"
    // Assume that credentials is a package that contains getters for any 
    // credentials needed (database dsn's etc.).
    //
    // For the purposes of the interview, you can use a sensible function 
    // name and assume it exists in this package (eg. credentials.MongoDbDsn())
    "umotif.com/go/credentials"
)

// The incoming event
type QuestionnaireCompletedEvent struct {
    Id string
    UserId string
    StudyId string
    QuestionnaireId string
    CompletedAt string
    RemainingCompletions int
}

func LambdaHander(ctx context.Context, event IncomingEvent) (string, error) {
    log.Print("Running ƛ %s", ctx.FunctionName)

    return fmt.Sprintf("Hello %s!", event.Name ), nil
}

func main() {
    lambda.Start(LambdaHander)
}
