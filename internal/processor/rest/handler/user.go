package handler

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	manager interfaces.IManager
}

func InitUserController(manager interfaces.IManager) *UserController {
	return &UserController{
		manager: manager,
	}
}

func (uc *UserController) PurchaseServiceByID(ctx *gin.Context) {

}
