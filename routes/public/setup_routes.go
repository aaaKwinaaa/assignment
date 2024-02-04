package public

import (
	"test/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, ctrl *app.SetupController) (*gin.RouterGroup, error) {
	api := router.Group("api/")
	{
		v1 := api.Group("v1")
		{
			boardCastRoute := v1.Group("boardCast")
			{
				SetupBoardCastRoutes(boardCastRoute, ctrl.BoardCastController)
			}
		}
	}

	return api, nil
}
