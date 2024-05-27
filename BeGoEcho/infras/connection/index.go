package connection

import (
	"database/sql"

	"github.com/CMDezz/KB/infras/logger"
	"github.com/jmoiron/sqlx"
)

func InitializeConnection(dbDriver string, dbUri string) (*sql.DB, *sqlx.DB) {

	//sql
	sqlDBContext, err := sql.Open(dbDriver, dbUri)
	if err != nil {
		logger.Error("Init connection Exception: %s", err)
		panic(err)
	}

	err = sqlDBContext.Ping()
	if err != nil {
		logger.Error("Init connection Exception: %s", err)
		panic(err)
	}

	//sqlx
	sqlxDBContext, err := sqlx.Open(dbDriver, dbUri)
	if err != nil {
		logger.Error("Init connection Exception: %s", err)
		panic(err)
	}

	err = sqlxDBContext.Ping()
	if err != nil {
		logger.Error("Init connection Exception: %s", err)
		panic(err)
	}

	return sqlDBContext, sqlxDBContext
}
