package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/CMDezz/KB/infras/token"
	"github.com/CMDezz/KB/utils/constants"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Custom middleware to check token expiration
func CheckTokenExpiration(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*token.Payload)

		// Check if token has expired
		if claims.ExpiresAt.Before(time.Now()) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token has expired")
		}
		return next(c)
	}
}

func CheckIsStaff(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.Payload)
	fmt.Println(claims)
	if claims.Role == constants.ENUM_PER_USER {
		return errors.New("permission denied")
	}

	return nil
}

func CheckIsAdmin(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.Payload)

	if claims.Role != constants.ENUM_PER_ADMIN {
		return errors.New("permission denied")
	}

	return nil

}
