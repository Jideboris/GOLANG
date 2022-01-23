package main

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"database/sql"
	// Assume that credentials is a package that contains getters for any
	// credentials needed (database dsn's etc.).
	//
	// For the purposes of the interview, you can use a sensible function
	// name and assume it exists in this package (eg. credentials.MongoDbDsn())
	"umotif.com/go/credentials"
)

// The incoming event
type QuestionnaireCompletedEvent struct {
	Id                   string
	UserId               string
	StudyId              string
	QuestionnaireId      string
	CompletedAt          string
	RemainingCompletions int
}

type ScheduledQuestionnaire struct {
	Id                   string
	UserId               string
	StudyId              string
	QuestionnaireId      string
}

type NewScheduleCreatedEvent struct {
	ScheduleId           string
}

type ScheduledUserQuestionnairesCompletedEvent struct {
	UserId               string
	CompletedAt			 string
}

func GenerateId() string {
	// returns uuid
}

func SaveToDatabase(messages chan<- NewScheduleCreatedEvent, event QuestionnaireCompletedEvent) {
	db, err := sql.Open("mysql", credentials.MySqlDsn())

	if err != nil {
		panic(err.Error())
	}

	newScheduledQuestionnaire := ScheduledQuestionnaire{
		Id: GenerateId(),
		UserId: event.UserId,
		StudyId: event.StudyId,
		QuestionnaireId: event.QuestionnaireId
	}

	// insert newScheduledQuestionnaire into database
	insert, err := db.Query("INSERT INTO `scheduled_questionnaires`...")

	if err != nil {
        panic(err.Error())
    }

	messages <- NewScheduleCreatedEvent{
		ScheduleId: newScheduledQuestionnaire.Id
	}

	// close db connection
	insert.Close()
}

func SendNewScheduleMessage(messages <-chan NewScheduleCreatedEvent) {
	// send message to sqs
}

func SendQuestionnairesCompletedMessage(messages <-chan ScheduledUserQuestionnairesCompletedEvent) {
	// send message to sqs
}

func LambdaHander(ctx context.Context, event QuestionnaireCompletedEvent) (string, error) {
	log.Print("Running ƛ %s", ctx.FunctionName)

	newSchedules = make(chan NewScheduleCreatedEvent)
	completedQuestionnaireUsers = make(chan ScheduledUserQuestionnairesCompletedEvent)

	if event.RemainingCompletions >= 1 {
		go SaveToDatabase(newSchedules, event)
	} else {
		completedQuestionnaireUsers <- ScheduledUserQuestionnairesCompletedEvent{
			UserId: event.UserId,
			CompletedAt: event.CompletedAt
		}
	}

	go SendNewScheduleMessage(newSchedules)
	go SendQuestionnairesCompletedMessage(completedQuestionnaireUsers)

	return fmt.Sprintf("Hello %s!", event.StudyId), nil
}

func main() {
	lambda.Start(LambdaHander)
}
