package apis

import (
	"database/sql"

	"github.com/CMDezz/KB/infras/apis/controllers"
	"github.com/CMDezz/KB/infras/apis/queries"
	"github.com/CMDezz/KB/infras/apis/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var Controller controllers.Controllers

func Initialize(e *echo.Echo, dbCtx *sql.DB, sqlxCtx *sqlx.DB) {
	//controller
	//service
	// Controller = &controllers.Controllers{}
	//queries

	queries := queries.NewQueries(dbCtx, sqlxCtx)
	Controller = controllers.Controllers{
		Services: services.NewServices(queries),
	}

	initRoutes(e)
}

func initRoutes(e *echo.Echo) {
	routes := e.Group("/apis/")
	// privateRoutes := e.Group("apis/priv/")

	//Public
	routes.POST("CreateDiscount", Controller.CreateDiscount)
	//Private
}
