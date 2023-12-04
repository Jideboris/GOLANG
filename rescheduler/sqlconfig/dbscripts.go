package sql_config

const (
	// create participant
	CreateParticipantData       string = "INSERT INTO `participants` (`id`,`name`) VALUES (?,?)"
	CreateParticipantTableQuery string = `
CREATE TABLE IF NOT EXISTS participants (id VARCHAR(128) PRIMARY KEY NOT NULL,name VARCHAR(128) NOT NULL)
`
	// create questionnaires
	CreateQuestionnairesData       string = "INSERT INTO `questionnaires` (`id`,`study_id`,`name`,`questions`,`max_attempts`,`hours_between_attempts`) VALUES (?,?,?,?,?,?)"
	CreateQuestionnairesTableQuery string = `
CREATE TABLE IF NOT EXISTS questionnaires (
	id VARCHAR(128) PRIMARY KEY NOT NULL,
	study_id VARCHAR(128) NOT NULL,
	name VARCHAR(128) NOT NULL,
	questions JSON NOT NULL,
	max_attempts INT,
	hours_between_attempts INT DEFAULT 24
);
`
	// create scheduled_questionnaires
	CreateScheduledQuestionnairesData       string = "INSERT INTO `scheduled_questionnaires` (`id`,`questionnaire_id`,`participant_id`,`scheduled_at`,`status`) VALUES (?,?,?,?,?)"
	CreateScheduledQuestionnairesTableQuery string = `
CREATE TABLE IF NOT EXISTS scheduled_questionnaires (
	id VARCHAR(128) PRIMARY KEY NOT NULL,
	questionnaire_id VARCHAR(128) NOT NULL,
	participant_id VARCHAR(128) NOT NULL,
	scheduled_at DATETIME NOT NULL,
	status VARCHAR(128) NOT NULL
);
`
	// create questionnaire_results
	CreateQuestionnairesResultsTableQuery string = `
CREATE TABLE IF NOT EXISTS questionnaire_results (
	id VARCHAR(128) NOT NULL,
	answers JSON NOT NULL,
	questionnaire_id VARCHAR(128) NOT NULL,
	participant_id VARCHAR(128) NOT NULL,
	questionnaire_schedule_id VARCHAR(128),
	completed_at DATETIME
);
`
)
