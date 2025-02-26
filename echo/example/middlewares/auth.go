package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Info("auth middleware")

		cookie, err := c.Cookie("auth")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if cookie.Value == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "unautharized")
		}

		c.Set("username", cookie.Value)
		return next(c)
	}
}
