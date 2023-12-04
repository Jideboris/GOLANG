package questionnaires_usecase

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"umotif.com/rescheduler/entity"
)

type usecase struct {
	tableQueries map[string]string
	dataQueries  map[string]entity.TupleData
	db           *sql.DB
	dbName       string
	ctx          context.Context
}

func NewUseCase(
	queries map[string]string,
	dataQ map[string]entity.TupleData,
	database *sql.DB,
	dbname string,
	ctx context.Context) UseCase {
	return &usecase{
		tableQueries: queries,
		dataQueries:  dataQ,
		db:           database,
		dbName:       dbname,
		ctx:          ctx,
	}
}

// CreateData implements UseCase.
func (uc *usecase) CreateData() (string, error) {
	var insertId string
	for k, query := range uc.dataQueries {
		if k == ParticipantTable {
			participant := entity.Participant{}
			participant.Id = query.Parameters.(*entity.Participant).Id
			participant.Name = query.Parameters.(*entity.Participant).Name
			_, err := uc.db.ExecContext(uc.ctx, query.Query, participant.Id, participant.Name)
			if err != nil {
				log.Fatalf("error inserting %s: %s", ParticipantTable, err)
				return "", fmt.Errorf("error inserting %s: %s", ParticipantTable, err)
			}
			insertId = participant.Id
		}

		if k == QuestionaireTable {
			questionaires := entity.Questionnaires{}
			questionaires.Id = query.Parameters.(*entity.Questionnaires).Id
			questionaires.Name = query.Parameters.(*entity.Questionnaires).Name
			questionaires.StudyId = query.Parameters.(*entity.Questionnaires).StudyId
			questionaires.Questions = query.Parameters.(*entity.Questionnaires).Questions
			questionaires.MaxAttempt = query.Parameters.(*entity.Questionnaires).MaxAttempt
			questionaires.HoursBetweenAttempts = query.Parameters.(*entity.Questionnaires).HoursBetweenAttempts

			_, err := uc.db.ExecContext(uc.ctx,
				query.Query,
				questionaires.Id,
				questionaires.StudyId,
				questionaires.Name,
				questionaires.Questions,
				questionaires.MaxAttempt,
				questionaires.HoursBetweenAttempts)
			if err != nil {
				log.Fatalf("error inserting %s: %s", QuestionaireTable, err)
				return "", fmt.Errorf("error inserting %s: %s", QuestionaireTable, err)
			}
			insertId = questionaires.Id
		}

		if k == ScheduledQuestionnairesTable {
			scheduledQuestionnaires := entity.ScheduledQuestionnaires{}
			scheduledQuestionnaires.Id = query.Parameters.(*entity.ScheduledQuestionnaires).Id
			scheduledQuestionnaires.ParticipantId = query.Parameters.(*entity.ScheduledQuestionnaires).ParticipantId
			scheduledQuestionnaires.QuestionnaireId = query.Parameters.(*entity.ScheduledQuestionnaires).QuestionnaireId
			scheduledQuestionnaires.ScheduledAt = query.Parameters.(*entity.ScheduledQuestionnaires).ScheduledAt
			scheduledQuestionnaires.Status = query.Parameters.(*entity.ScheduledQuestionnaires).Status
			_, err := uc.db.ExecContext(uc.ctx,
				query.Query,
				scheduledQuestionnaires.Id,
				scheduledQuestionnaires.QuestionnaireId,
				scheduledQuestionnaires.ParticipantId,
				scheduledQuestionnaires.ScheduledAt,
				scheduledQuestionnaires.Status)
			if err != nil {
				log.Fatalf("error inserting participants: %s", err)
				return "", fmt.Errorf("error inserting participants: %s", err)
			}
		}
	}

	return insertId, nil
}

// CreateTable implements UseCase.
func (uc *usecase) CreateTable() (bool, error) {
	for _, query := range uc.tableQueries {
		res, err := uc.db.ExecContext(uc.ctx, query)
		if err != nil {
			log.Printf("Error %s when executing query\n", err)
			return false, fmt.Errorf("error %s when executing query", err)
		}
		no, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when fetching rows", err)
			return false, fmt.Errorf("error %s when fetching rows", err)
		}
		log.Printf("rows affected: %d\n", no)
	}
	return true, nil

}
