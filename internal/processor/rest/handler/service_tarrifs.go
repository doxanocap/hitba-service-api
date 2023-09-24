package handler

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/errs"
	"github.com/doxanocap/pkg/lg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServiceTariffsController struct {
	manager interfaces.IManager
}

func InitServiceTariffsController(manager interfaces.IManager) *ServiceTariffsController {
	return &ServiceTariffsController{
		manager: manager,
	}
}

func (sc *ServiceTariffsController) Create(ctx *gin.Context) {
	var tariff model.ServiceTariff

	if err := ctx.ShouldBindJSON(&tariff); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := sc.manager.Service().ServiceTariffs().Create(ctx, tariff)
	if err != nil {
		lg.Errorf("service_tariffs_ctl.Create: %v", err)

		code := errs.UnmarshalCode(err)
		if code == http.StatusNotFound {
			ctx.JSON(code, model.ErrServiceIdNotFound)
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
		return
	}
}
