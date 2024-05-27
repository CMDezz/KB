package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CMDezz/KB/infras/connection"
	"github.com/CMDezz/KB/infras/logger"
	"github.com/CMDezz/KB/utils"
	"github.com/CMDezz/KB/utils/constants"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	/* ---------------------------------- */
	/*      //load application config     */
	/* ---------------------------------- */

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
	}

	/* ---------------------------------- */
	/*            //init logger           */
	/* ---------------------------------- */
	logger.NewLogger(config.LoggerPath, config.Enviroment)

	/* ---------------------------------- */
	/*         //init DB connections         */
	/* ---------------------------------- */
	sqlDBContext, sqlxDBContext := connection.InitializeConnection(config.DbDriver, config.DbSource)
	//close DB connections on defer
	defer func(sqlDBContext *sql.DB) {
		sqlDBContext.Close()
	}(sqlDBContext)

	defer func(sqlxDBContext *sqlx.DB) {
		sqlxDBContext.Close()
	}(sqlxDBContext)

	fmt.Println("Connect to db succesfully")

	/* ---------------------------------- */
	/*         //init echo                */
	/* ---------------------------------- */
	e := echo.New()
	e.Server.SetKeepAlivesEnabled(false)
	e.Server.ReadTimeout = constants.TimeoutServerDefault
	e.Server.WriteTimeout = constants.TimeoutServerDefault

	e.Use(middleware.CORS())

	//init echo routes

}
