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

type UserController struct {
	manager interfaces.IManager
	cfg     *model.Config
	log     *zap.Logger
}

func InitUserController(manager interfaces.IManager, cfg *model.Config) *UserController {
	return &UserController{
		manager: manager,
		cfg:     cfg,
		log:     logger.Log.Named("[CONTROLLER][USER]"),
	}
}

func (ctl *UserController) PurchaseServiceByID(ctx *gin.Context) {
	log := ctl.log.Named("[PurchaseServiceByID]")

	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		log.Error(fmt.Sprintf("parsing param: %s", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
	}

	var purchase model.Purchase
	if err := ctx.ShouldBindJSON(&purchase); err != nil {
		log.Error(fmt.Sprintf("bindJSON: %s", err))
		ctx.JSON(http.StatusBadRequest, model.HttpBadRequest)
		return
	}

	err = ctl.manager.Service().User().PurchaseService(ctx, &purchase, int64(userID))
	if err != nil {
		log.Error(fmt.Sprintf("service.user.PurchaseService: %v", err))

		code := errs.UnmarshalCode(err)
		if code == http.StatusNotFound {
			ctx.JSON(code, model.ErrServiceIdNotFound)
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.HttpInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}
