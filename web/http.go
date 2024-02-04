package web

import (
	"test/app"
	"test/infrastructure"
	"test/infrastructure/handler"
	publicRoutes "test/routes/public"
)

func StartServer(httpHandler handler.HTTP, app *app.App, adt infrastructure.Infrastructure) {
	if _, err := publicRoutes.SetupRoutes(&httpHandler.Connection.RouterGroup, &app.Controller); err != nil {
		panic(err)
	}

	if err := httpHandler.Connection.Run(); err != nil {
		panic("Failed to start the project")
	}
}
