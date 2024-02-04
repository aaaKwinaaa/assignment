package app

import "test/service"

type SetupService struct {
	Integration *SetupIntegration
	BoardCastService       service.BoardCastService
}

func (i *SetupService) Setup() {
	i.BoardCastService = service.NewBoardCastService(i.Integration.TransactionIntegration)

}
