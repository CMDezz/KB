package apis

import (
	"database/sql"
	"log"

	"github.com/CMDezz/KB/infras/apis/controllers"
	"github.com/CMDezz/KB/infras/apis/queries"
	"github.com/CMDezz/KB/infras/apis/services"
	"github.com/CMDezz/KB/infras/middleware"
	"github.com/CMDezz/KB/infras/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var Controller controllers.Controllers

func Initialize(e *echo.Echo, dbCtx *sql.DB, sqlxCtx *sqlx.DB, secretKey string) {
	//controller
	//service
	// Controller = &controllers.Controllers{}
	//queries

	tokenMarker, err := token.NewJWTTokenMaker(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	queries := queries.NewQueries(dbCtx, sqlxCtx)
	Controller = controllers.Controllers{
		Services: services.NewServices(queries, tokenMarker),
		Token:    *tokenMarker,
	}

	initRoutes(e, secretKey)

}

func initRoutes(e *echo.Echo, secretKey string) {
	routes := e.Group("/apis/")
	// token := &token.Payload{}
	// middleware.AuthMiddleware(&Controller.Token)
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.Payload)
		},
		SigningKey: []byte(secretKey),
	}
	privateRoutes := e.Group("apis/")
	privateRoutes.Use(echojwt.WithConfig(config))
	privateRoutes.Use(middleware.CheckTokenExpiration)

	//Public
	routes.POST("CreateDiscount", Controller.CreateDiscount)
	routes.POST("CreateAccount", Controller.CreateAccount)
	routes.GET("GetAllDiscount", Controller.GetAllDiscount)
	routes.GET("GetAllAccount", Controller.GetAllAccount)
	routes.GET("GetDiscountById/:id", Controller.GetDiscountById)
	routes.PUT("UpdateDiscountById", Controller.UpdateDiscountById)
	routes.POST("LoginAccount", Controller.LoginAccount)
	//Private
	privateRoutes.DELETE("DeleteDiscountById/:id", Controller.DeleteDiscountById)
}
