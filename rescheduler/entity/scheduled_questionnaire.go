package entity

import "time"

// A scheduled_questionnaires Struct to map every scheduled_questionnaires
type ScheduledQuestionnaires struct {
	Id              string    `json:"id"`
	QuestionnaireId string    `json:"questionnaire_id"`
	ParticipantId   string    `json:"participant_id"`
	ScheduledAt     time.Time `json:"scheduled_at"`
	Status          string    `json:"status"`
}