package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBHost         string `envconfig:"DB_HOST" default:"localhost:3306"`
	DBPort         string `envconfig:"DB_PORT" default:"3306"`
	DBUserName     string `envconfig:"DB_USERNAME" default:"gotest"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"password"`
	DBDatabaseName string `envconfig:"DB_DBNAME" default:"questionnairesDB"`
	SQSqueue       string `envconfig:"SQS_QUEUE" default:"QUEUE URL"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	cfg.DBPassword = "1234"
	cfg.DBUserName = "gotest"
	cfg.DBHost = "localhost:3306"
	cfg.DBDatabaseName = "questionnairesDB"
	return cfg
}
