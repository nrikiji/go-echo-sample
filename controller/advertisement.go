package controller

import(
	"bytes"
	"net/http"
	"text/template"
	
	"github.com/labstack/echo"
)

type Advertise struct {
	Id int `json:"id"`
	Html string `json:"html"`
}

type TemplateData struct {
	Name string
	Message string
}

func GetAdvertisement() echo.HandlerFunc {
	return func(c echo.Context) error {

		data := TemplateData{"タイトル", "本文"}

		var buffer bytes.Buffer
		var html = "<div><h1>{{.Name}}</h1><p>{{.Message}}</p></div>"
		var t = template.Must(template.New("html").Parse(html))

		if err := t.Execute(&buffer, data); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
		} else {
			return c.JSON(http.StatusOK, Advertise {
				Id: 1,
				Html: buffer.String(),
			})
		}
	}
}
