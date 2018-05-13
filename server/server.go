package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/piyush2206/go-to-do/app"
)

// Start initialises and starts HTTP server
func Start(appCtx *app.Ctx) {
	appServer := echo.New()
	appServer.HideBanner = true

	attachAPIs(appServer)

	appServer.Logger.Fatal(appServer.Start(":1323"))
}

// attachAPIs adds API endpoints to the HTTP server
func attachAPIs(appServer *echo.Echo) {
	appServer.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hullo!")
	})
}
