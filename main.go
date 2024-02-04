package main

import (
	"test/app"
	"test/infrastructure"
	"test/web"
)

func main() {

	adt := infrastructure.NewInfrastructure()
	app := app.New(adt)

	web.StartServer(adt.HTTP, app, adt)

}
