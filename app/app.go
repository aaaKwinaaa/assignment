package app

import (
	"test/infrastructure"
)

type App struct {
	Integration SetupIntegration
	Controller  SetupController
	Service     SetupService
}

func New(adt infrastructure.Infrastructure) *App {
	integration := &SetupIntegration{}
	integration.Setup()

	service := &SetupService{Integration: integration}
	service.Setup()

	controller := &SetupController{Service: service}
	controller.Setup()

	return &App{
		Integration: *integration,
		Controller:  *controller,
		Service:     *service,
	}
}
