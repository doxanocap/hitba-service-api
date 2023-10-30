package controllers

import (
	"fmt"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/hitba-service-api/pkg/logger"
	"github.com/doxanocap/pkg/errs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ServicesController struct {
	manager interfaces.IManager
	config  *model.Config
	log     *zap.Logger
}

func InitServicesController(manager interfaces.IManager, config *model.Config) *ServicesController {
	return &ServicesController{
		manager: manager,
		config:  config,
		log:     logger.Log.Named("[CONTROLLER][SERVICE]"),
	}
}

// Create
//
//	@Summary		createServices
//	@Tags			services
//	@Description	create service
//	@Accept			json
//	@Produce		json
//	@Param			service	body	model.Service	true	"service"
//	@Success		200
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		409	{object}	model.ErrorResponse
//	@Failure		500	{object}	model.ErrorResponse
//	@Failure		default
//	@Router			/v1/services [post]
func (ctl *ServicesController) Create(ctx *gin.Context) {
	log := ctl.log.Named("[Create]")

	var service model.Service
	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Error(fmt.Sprintf("bindJSON: %s", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	err := ctl.manager.Service().Services().Create(ctx, &service)
	if err != nil {
		log.Error(fmt.Sprintf("service.services.Create: %v", err))

		code := errs.UnmarshalCode(err)
		switch code {
		case http.StatusConflict:
			ctx.JSON(code, model.ErrSuchServiceAlreadyExist)
		default:
			ctx.JSON(code, model.HttpInternalServerError)
		}
		return
	}

	ctx.Status(http.StatusOK)
}

// CreateTariff
//
//	@Summary		createTariff
//	@Tags			services
//	@Description	create service_tariff using service id
//	@Accept			json
//	@Produce		json
//	@Param			tariff	body	model.ServiceTariff	true	"service_tariff"
//	@Success		200
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		404	{object}	model.ErrorResponse
//	@Failure		500	{object}	model.ErrorResponse
//	@Failure		default
//	@Router			/v1/services/tariffs [post]
func (ctl *ServicesController) CreateTariff(ctx *gin.Context) {
	log := ctl.log.Named("[CreateTariff]")

	var tariff model.ServiceTariff
	if err := ctx.ShouldBindJSON(&tariff); err != nil {
		log.Error(fmt.Sprintf("bindJSON: %s", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	err := ctl.manager.Service().Services().CreateTariff(ctx, &tariff)
	if err != nil {
		code := errs.UnmarshalCode(err)
		log.Error(fmt.Sprintf("service.services.CreateTariff: %s", errs.UnmarshalMsg(err)))
		switch code {
		case http.StatusNotFound:
			ctx.JSON(code, err)
		default:
			ctx.JSON(code, model.HttpInternalServerError)
		}
		return
	}

	ctx.Status(http.StatusOK)
}

// GetAll
//
//	@Summary		getAll
//	@Tags			services
//	@Description	get all info about services
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		500	{object}	model.ErrorResponse
//	@Failure		default
//	@Router			/v1/services [get]
func (ctl *ServicesController) GetAll(ctx *gin.Context) {
	log := ctl.log.Named("[GetAll]")

	result, err := ctl.manager.Service().Services().GetAllServices(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("service.services.GetAll: %s", err))
		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
	}

	ctx.JSON(http.StatusOK, result)
}

// UpdateTariff
//
//	@Summary		updateTariff
//	@Tags			services
//	@Description	update service_tariff using service_id and service_tariff_id
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string				true	"service_tariff_id"
//	@Param			tariff	body	model.ServiceTariff	true	"service_tariff"
//	@Success		200
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		404	{object}	model.ErrorResponse
//	@Failure		500	{object}	model.ErrorResponse
//	@Failure		default
//	@Router			/v1/services/tariffs [put]
func (ctl *ServicesController) UpdateTariff(ctx *gin.Context) {
	log := ctl.log.Named("[UpdateTariff]")

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Error(fmt.Sprintf("parsing param: %s", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
	}

	var tariff model.ServiceTariff
	if err := ctx.ShouldBindJSON(&tariff); err != nil {
		log.Error(fmt.Sprintf("bindJSON: %s", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	err = ctl.manager.Service().Services().UpdateTariffByID(ctx, &tariff, int64(id))
	if err != nil {
		log.Error(fmt.Sprintf("service.services.UpdateTariffByID: %s", err))
		code := errs.UnmarshalCode(err)
		switch code {
		case http.StatusNotFound:
			ctx.JSON(code, err)
		default:
			ctx.JSON(code, model.HttpInternalServerError)
		}
		return
	}
	ctx.Status(http.StatusOK)
}
