package questionnaires_usecase

type UseCase interface {
	CreateTable() (bool, error)
	CreateData() (string, error)
}

// Tables enum
const (
	ParticipantTable             string = "participant"
	QuestionaireTable            string = "questionaires"
	ScheduledQuestionnairesTable string = "scheduled_questionnaires"
	QuestionnairesResults        string = "questionaires_results"
)

// Question status enum
const (
	Pending   string = "pending"
	Completed string = "completed"
)
