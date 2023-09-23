package handler

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServicesController struct {
	manager interfaces.IManager
}

func InitServicesController(manager interfaces.IManager) *ServicesController {
	return &ServicesController{
		manager: manager,
	}
}

func (sc *ServicesController) Create(ctx *gin.Context) {

}

func (sc *ServicesController) GetAll(ctx *gin.Context) {
	result := sc.manager.Service().Services().GetAll(ctx)
	//if err != nil {
	//	if errs.IsHttpNotFoundError(err) {
	//		ctx.Status(http.StatusNotFound)
	//		return
	//	}
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//}

	ctx.JSON(http.StatusOK, result)
}
