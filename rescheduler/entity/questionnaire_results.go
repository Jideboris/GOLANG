package entity

import "time"

// A questionnaire_results Struct to map every questionnaire_results
type QuestionnaireResults struct {
	Id                      string    `json:"id"`
	Answers                 string    `json:"answers"`
	QuestionnaireId         string    `json:"questionnaire_id"`
	ParticipantId           string    `json:"participant_id"`
	QuestionnaireScheduleId string    `json:"questionnaire_schedule_id"`
	CompletedAt             time.Time `json:"completed_at"`
}
