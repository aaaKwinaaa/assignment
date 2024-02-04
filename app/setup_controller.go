package app

import controller "test/controller/public"

type SetupController struct {
	Service             *SetupService
	BoardCastController controller.BoardCastController
}

func (i *SetupController) Setup() {
	i.BoardCastController = controller.NewBoardCastController(i.Service.BoardCastService)
}
