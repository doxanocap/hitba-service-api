package controllers

import (
	"fmt"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/hitba-service-api/pkg/logger"
	"github.com/doxanocap/pkg/errs"
	"github.com/doxanocap/pkg/lg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type ServicesController struct {
	manager interfaces.IManager
	cfg     *model.Config
	log     *zap.Logger
}

func InitServicesController(manager interfaces.IManager, cfg *model.Config) *ServicesController {
	return &ServicesController{
		manager: manager,
		cfg:     cfg,
		log:     logger.Log.Named("[CONTROLLER][SERVICE]"),
	}
}

func (ctl *ServicesController) Create(ctx *gin.Context) {
	log := ctl.log.Named("[Create]")

	var service model.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Error(fmt.Sprintf("bindJSON: %s", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	err := ctl.manager.Service().Services().Create(ctx, service)
	if err != nil {
		lg.Errorf("service.services.Create: %v", err)

		code := errs.UnmarshalCode(err)
		if code == http.StatusConflict {
			ctx.JSON(code, model.ErrSuchServiceAlreadyExist)
			return
		}
		ctx.JSON(code, model.HttpInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (ctl *ServicesController) CreateTariff(ctx *gin.Context) {
	log := ctl.log.Named("[Create]")

	var tariff model.ServiceTariff

	if err := ctx.ShouldBindJSON(&tariff); err != nil {
		log.Error(fmt.Sprintf("bindJSON: %s", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	err := ctl.manager.Service().Services().CreateTariff(ctx, tariff)
	if err != nil {
		log.Error(fmt.Sprintf("service.serviceTariffs.Create: %s", err))

		code := errs.UnmarshalCode(err)
		if code == http.StatusNotFound {
			ctx.JSON(code, model.ErrServiceIdNotFound)
			return
		}

		ctx.JSON(code, model.HttpInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (ctl *ServicesController) GetAll(ctx *gin.Context) {
	log := ctl.log.Named("[GetAll]")

	result, err := ctl.manager.Service().Services().GetAllServices(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("service.services.GetAll: %s", err))
		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
	}

	ctx.JSON(http.StatusOK, result)
}
