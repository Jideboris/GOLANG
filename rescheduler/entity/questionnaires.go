package entity

// A questionnaires Struct to map every Questionnaire
type Questionnaires struct {
	Id                   string `json:"id"`
	StudyId              string `json:"study_id"`
	Name                 string `json:"name"`
	Questions            string `json:"questions"`
	MaxAttempt           int64  `json:"max_attempts"`
	HoursBetweenAttempts int64  `json:"hours_between_attempts"`
}
