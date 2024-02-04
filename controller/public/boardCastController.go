package controller

import (
	"net/http"
	"test/handler"
	"test/service"
	"test/service/dto"

	"github.com/gin-gonic/gin"
)

type BoardCastController interface {
	BoardCastTransaction(c *gin.Context)
	UtilizeTransaction(c *gin.Context)
}

type boardCastController struct {
	boardCastService service.BoardCastService
}

func NewBoardCastController(boardCastService service.BoardCastService) BoardCastController {
	return &boardCastController{
		boardCastService: boardCastService,
	}
}

func (controller *boardCastController) BoardCastTransaction(c *gin.Context) {

	var request dto.BoardCastRequestDto
	if err := c.BindJSON(&request); err != nil {
		return
	}

	data, err := controller.boardCastService.BoardCastTransaction(c, request)
	if err != nil {
		_ = c.Error(err)
		panic(err)
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       data,
	})
}

func (controller *boardCastController) UtilizeTransaction(c *gin.Context) {

	data, err := controller.boardCastService.UtilizeTransaction(c, c.Param("hash"))
	if err != nil {
		_ = c.Error(err)
		panic(err)
	}

	c.JSON(http.StatusOK, handler.Wrapper{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       data,
	})
}
