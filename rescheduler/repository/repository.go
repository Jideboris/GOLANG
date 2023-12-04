package questionnaires_repository

import (
	"context"
	"database/sql"

	"umotif.com/rescheduler/app/repository"
)

func New(db *sql.DB, ctx context.Context) repository.DbRepositoryParams {
	return repository.DbRepositoryParams{
		Db:  db,
		Ctx: ctx}
}

func NewQuery(query string) repository.DbQueryParams {
	return repository.DbQueryParams{
		Query: query,
	}
}
