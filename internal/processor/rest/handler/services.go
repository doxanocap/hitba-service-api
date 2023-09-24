package handler

import (
	"fmt"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/hitba-service-api/pkg/logger"
	"github.com/doxanocap/pkg/errs"
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
	var service model.Service

	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := sc.manager.Service().Services().Create(ctx, service)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("servicesController.Create: %v", err))

		code := errs.UnmarshalCode(err)
		if code == http.StatusConflict {
			ctx.JSON(code, model.ErrSuchServiceAlreadyExist)
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
		return
	}
}

func (sc *ServicesController) GetAll(ctx *gin.Context) {
	result, err := sc.manager.Service().Services().GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
	}

	ctx.JSON(http.StatusOK, result)
}
