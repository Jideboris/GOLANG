package main

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"database/sql"
	"github.com/nu7hatch/gouuid"
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
	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err.Error())
	}
	return uuid.String()
}

func SaveToDatabase(event QuestionnaireCompletedEvent) ScheduledQuestionnaire {
	// open db connection with known credentials
	db, err := sql.Open("mysql", credentials.MySqlDsn())

	if err != nil {
		panic(err.Error())
	}

	newScheduledQuestionnaire := ScheduledQuestionnaire{
		Id: GenerateId(),
		UserId: event.UserId,
		StudyId: event.StudyId,
		QuestionnaireId: event.QuestionnaireId,
	}

	// insert newScheduledQuestionnaire into database
	insert, err := db.Query("INSERT INTO `scheduled_questionnaires`...")

	if err != nil {
        panic(err.Error())
    }

	log.Printf("Schedule %s created.", newScheduledQuestionnaire.Id)

	// close db connection
	insert.Close()

	return newScheduledQuestionnaire
}

func SendNewScheduleMessage(questionnaire ScheduledQuestionnaire) {
	message := NewScheduleCreatedEvent{
		ScheduleId: questionnaire.Id,
	}

	// dispatch message to sqs...

	log.Printf("Schedule %s creation message sent.", message.ScheduleId)
}

func SendQuestionnairesCompletedMessage(event QuestionnaireCompletedEvent) {
	message := ScheduledUserQuestionnairesCompletedEvent{
		UserId: event.UserId,
		CompletedAt: event.CompletedAt,
	}

	// dispatch message to sqs...

	log.Printf("All questionnaires for user %s completed.", message.UserId)
}

func LambdaHander(ctx context.Context, event QuestionnaireCompletedEvent) (string, error) {
	log.Print("Running ƛ %s", ctx.FunctionName)

	if event.RemainingCompletions >= 1 {
		scheduledQuestionnaire := SaveToDatabase(event)
		SendNewScheduleMessage(scheduledQuestionnaire)
	} else {
		SendQuestionnairesCompletedMessage(event)
	}

	return fmt.Sprintf("Hello %s!", event.StudyId), nil
}

func main() {
	lambda.Start(LambdaHander)
}
