package controller

import(
	"app/model"
	"net/http"
	
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetTraking() echo.HandlerFunc {
	return func(c echo.Context) error {

		var campaign model.Campaign
		db := model.GetConnection()
		data := db.First(&campaign, 1)

		if data.Error == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Not Found")
		} else {
			return c.JSON(http.StatusOK, data)
		}
	}
}
