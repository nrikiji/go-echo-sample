package controller

import(
	MyErr "app/error"
	"net/http"
	// "errors"
	
	"github.com/labstack/echo"
)

type Ad struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func GetTraking() echo.HandlerFunc {
	return func(c echo.Context) error {

		ads := []Ad{{
			Id: 1,
			Name: "test1",
		}, {
			Id: 2,
			Name: "test2",
		}}

		return echo.NewHTTPError(http.StatusInternalServerError, "test")

		// エラーを投げる
		//return errors.New("MyError")
		return &MyErr.SystemError{Message: "エラーだよ"}

		// テキストを返す
		// return c.String(http.StatusOK, "traking")

		// JSONを返す
		return c.JSON(http.StatusOK, ads)
		
	}
}
