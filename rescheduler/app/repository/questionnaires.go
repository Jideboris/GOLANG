package repository

import (
	"context"
	"database/sql"
)

type QuestionnairesRepository interface {
	CreateTable() (bool, error)
	CreateData(parameters interface{}, table string) (string, error)
}

type DbRepositoryParams struct {
	Db    *sql.DB
	Ctx   context.Context
}

type DbQueryParams struct {
	Query string
}