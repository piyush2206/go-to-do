package server

import (
	"github.com/labstack/echo"
	"github.com/piyush2206/go-to-do/app"
	"github.com/piyush2206/go-to-do/controllers"
)

// Start initialises and starts HTTP server
func Start(appCtx *app.Context) {
	server := echo.New()
	server.HideBanner = true

	attachAPIs(server, appCtx)

	server.Logger.Fatal(server.Start(":1323"))
}

// attachAPIs adds API routes to the HTTP server
func attachAPIs(server *echo.Echo, appCtx *app.Context) {
	server.Use(bindAppCtx(appCtx))

	server.GET("/lists", controllers.GetLists)
	server.GET("/lists/:id", controllers.GetLists)
	server.POST("/lists", controllers.CreateList)
	server.DELETE("/lists/:id", controllers.DeleteList)

	server.POST("/tasks", controllers.CreateTask)
	server.PUT("/tasks/:id", controllers.UpdateTask)
	server.DELETE("/tasks/:id", controllers.DeleteTask)
}

// bindAppCtx sets app context to api context object for dependency injection
func bindAppCtx(appCtx *app.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("appCtx", appCtx)
			return next(c)
		}
	}
}
