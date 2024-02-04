package public

import (
	controller "test/controller/public"

	"github.com/gin-gonic/gin"
)

func SetupBoardCastRoutes(r *gin.RouterGroup, ctrl controller.BoardCastController) {
	r.POST("transaction", ctrl.BoardCastTransaction)
	r.GET("transaction/:hash" , ctrl.UtilizeTransaction)

}
