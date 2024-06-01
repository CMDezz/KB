package queries

import (
	"context"
	"database/sql"

	"github.com/CMDezz/KB/dto"
	"github.com/CMDezz/KB/infras/logger"
	"github.com/jmoiron/sqlx"
)

type PostgresQueries struct {
	DbContext     *sql.DB
	SQLxDBContext *sqlx.DB
}

type Queries struct {
	Postgres PostgresQueries
}

func NewQueries(dbCtx *sql.DB, SQLxDBContext *sqlx.DB) IQueries {
	queries := Queries{}
	queries.Postgres.SetDbContext(dbCtx, SQLxDBContext)
	return &queries
}
func (postgres *PostgresQueries) SetDbContext(dbContext *sql.DB, sqlxDbContext *sqlx.DB) {
	postgres.DbContext = dbContext
	postgres.SQLxDBContext = sqlxDbContext
}

func (postgres *PostgresQueries) HandleError(err error, query string) {
	logger.Error("[POSTGRES] - Error: %s, Query: %s", err, query)
}

type IQueries interface {
	DBCreateDiscount(ctx context.Context, discount *dto.Discount) (*dto.Discount, error)
}
