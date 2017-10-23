package route

import(
	"github.com/labstack/echo"

	"app/controller"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/tracking", controller.GetTraking())

	return e
}
