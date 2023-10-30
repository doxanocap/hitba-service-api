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
	config  *model.Config
	log     *zap.Logger
}

func InitUserController(manager interfaces.IManager, config *model.Config) *UserController {
	return &UserController{
		manager: manager,
		config:  config,
		log:     logger.Log.Named("[CONTROLLER][USER]"),
	}
}

// PurchaseServiceByID
//	@Summary		purchaseServiceByID
//	@Tags			users
//	@Description	purchase service by service_tariff id
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path	string	true	"user_id"
//	@Success		200
//	@Failure		400	{object}	model.ErrorResponse
//	@Failure		404	{object}	model.ErrorResponse
//	@Failure		500	{object}	model.ErrorResponse
//	@Failure		default
//	@Router			/v1/users/:user_id [post]
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
		switch code {
		case http.StatusNotFound:
			ctx.JSON(code, model.ErrServiceIDNotFound)
		default:
			ctx.JSON(code, model.HttpInternalServerError)
		}
		return
	}
	ctx.Status(http.StatusOK)
}
