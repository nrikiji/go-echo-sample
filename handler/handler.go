package handler

import(
	AppErr "app/error"
	"net/http"
	"github.com/labstack/echo"
)

type ApiError struct {
	Status int `json:status`
	Message string `json:message`
}

func JSONErrorHandler(err error, c echo.Context) {

	switch e := err.(type) {
	case *AppErr.BusinessError:
		// Business Error
		c.JSON(
			http.StatusOK, ApiError{
			Status: 200,
			Message: e.Message,
		})
	case *AppErr.SystemError:
		// System Error
		c.JSON(http.StatusOK, ApiError{
			Status: 500,
			Message: e.Message,
		})
	default:
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				// 404
				c.JSON(he.Code, ApiError{
					Status: he.Code,
					Message: "Not Found",
				})
			} else {
				// その他サーバーエラー
				c.JSON(he.Code, ApiError{
					Status: he.Code,
					Message: "System Error",
				})
			}
		}
	}
	
}
