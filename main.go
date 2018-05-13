package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	appServer := echo.New()
	appServer.HideBanner = true

	appServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hullo!")
	})
	appServer.Logger.Fatal(appServer.Start(":1323"))
}
