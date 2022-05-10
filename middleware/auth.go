package middleware

import (
	"net/http"
	"strings"

	"go-echo-starter/store"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(as store.AuthStore, ds store.DataStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			reqToken := c.Request().Header.Get(echo.HeaderAuthorization)
			splitToken := strings.Split(reqToken, "Bearer: ")

			if len(splitToken) < 2 {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			uid, err := as.VerifyIdToken(splitToken[1])
			if err != nil {
				return c.JSON(http.StatusUnauthorized, nil)
			}

			user, err := ds.FindUserByUid(uid)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, nil)
			} else if user == nil {
				return c.JSON(http.StatusForbidden, nil)
			}

			c.Set("user", user)

			return next(c)
		}
	}
}
