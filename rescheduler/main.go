package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"umotif.com/rescheduler/app/usecase/questionnaires_usecase"
	"umotif.com/rescheduler/config"
	"umotif.com/rescheduler/entity"
	"umotif.com/rescheduler/handler"
	sql_config "umotif.com/rescheduler/sqlconfig"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	// Assume that credentials is a package that contains getters for any
	// credentials needed (database dsn's etc.).
	//
	// For the purposes of the interview, you can use a sensible function
	// name and assume it exists in this package (eg. credentials.MongoDbDsn())
	// "umotif.com/go/credentials"
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

func SendLambdaMessage(ctx context.Context, event QuestionnaireCompletedEvent, wg *sync.WaitGroup) {
	defer wg.Done()
	lambda.Start(LambdaHander)
}

func LambdaHander(ctx context.Context, event QuestionnaireCompletedEvent) (string, error) {
	session, _ := session.NewSession()
	ptr := "PATH" //ctx.Value("QUEUE_URL").(string)
	newSqs := sqs.New(session)
	body, _ := json.Marshal(event)
	messageGroupId := "message1"
	_, err := newSqs.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:            new(int64),
		MessageAttributes:       map[string]*sqs.MessageAttributeValue{},
		MessageBody:             aws.String(string(body)),
		MessageDeduplicationId:  new(string),
		MessageGroupId:          &messageGroupId,
		MessageSystemAttributes: map[string]*sqs.MessageSystemAttributeValue{},
		QueueUrl:                &ptr,
	})
	return "", err
}

func dsn(dbName string, con *sql_config.Sqlconfiguration) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", con.DBUserName, con.DBPass, con.DBHost, dbName)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	cfg := config.New()

	configuration := sql_config.Connect(cfg.DBHost,
		cfg.DBPort,
		cfg.DBUserName,
		cfg.DBPassword,
		cfg.DBDatabaseName)

	db, err := sql.Open("mysql", dsn("", configuration))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+configuration.DBDatabaseName)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected: %d\n", no)
	defer db.Close()

	db, err = sql.Open("mysql", dsn(configuration.DBDatabaseName, configuration))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}

	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}

	go startUp(db, cfg, &wg)

    go SendLambdaMessage(ctx, QuestionnaireCompletedEvent{}, &wg)

	wg.Wait()
}

func startUp(db *sql.DB, cfg config.Config, wg *sync.WaitGroup) {
	defer wg.Done()

	now := time.Now()
	// queries to create required tables
	queries := make(map[string]string)
	queries[questionnaires_usecase.ParticipantTable] = sql_config.CreateParticipantTableQuery
	queries[questionnaires_usecase.QuestionaireTable] = sql_config.CreateQuestionnairesTableQuery
	queries[questionnaires_usecase.ScheduledQuestionnairesTable] = sql_config.CreateScheduledQuestionnairesTableQuery
	queries[questionnaires_usecase.QuestionnairesResults] = sql_config.CreateQuestionnairesResultsTableQuery

	// create unigue ID for each row of data inserted into tables
	uuidQuestionnaries := uuid.Must(uuid.NewRandom()).String()
	uuidParticipant := uuid.Must(uuid.NewRandom()).String()
	uuidScheduledQuestionnaires := uuid.Must(uuid.NewRandom()).String()

	// create json questionnaire questions to be inserted into table.
	dataQuestions := map[string]interface{}{
		"question": "What is your name?",
	}
	jsonStr, _ := json.Marshal(dataQuestions)

	// create mocked data generation querries and data for each table.
	dataCreationData := make(map[string]entity.TupleData)
	dataCreationData[questionnaires_usecase.ParticipantTable] = entity.TupleData{
		Query: sql_config.CreateParticipantData, Parameters: &entity.Participant{Id: uuidParticipant, Name: "Anthony Borisade"}}
	dataCreationData[questionnaires_usecase.QuestionaireTable] = entity.TupleData{
		Query: sql_config.CreateQuestionnairesData, Parameters: &entity.Questionnaires{Id: uuidQuestionnaries, StudyId: "Study 1", Name: "Eco system", Questions: string(jsonStr), MaxAttempt: 10, HoursBetweenAttempts: 2}}
	dataCreationData[questionnaires_usecase.ScheduledQuestionnairesTable] = entity.TupleData{
		Query: sql_config.CreateScheduledQuestionnairesData, Parameters: &entity.ScheduledQuestionnaires{Id: uuidScheduledQuestionnaires,
			QuestionnaireId: uuidQuestionnaries,
			ParticipantId:   uuidParticipant,
			ScheduledAt:     now,
			Status:          questionnaires_usecase.Completed}}

	// set the background context to hold the QueryURL for the SQS queue
	c := context.Background()
	//TODO : fix use built in string
	c2 := context.WithValue(c, "QUEUE_URL", cfg.SQSqueue)

	repo_uc := questionnaires_usecase.NewUseCase(queries, dataCreationData,
		db, cfg.DBDatabaseName, c2)

	_, err := repo_uc.CreateTable()
	if err != nil {
		log.Panicln("Failed to create required tables:", err)
		return
	}

	_, err = repo_uc.CreateData()
	if err != nil {
		log.Panicln("Failed to insert required tables data:", err)
		return
	}

	handler.NewHandler(repo_uc)

}
