package handler

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest/controllers"
	gin2 "github.com/doxanocap/hitba-service-api/pkg/gin"
	"github.com/gin-gonic/gin"
	"sync"
)

type Handler struct {
	services       *controllers.ServicesController
	serviceTariffs *controllers.ServiceTariffsController
	user           *controllers.UserController

	engine       *gin.Engine
	engineRunner sync.Once
	manager      interfaces.IManager
}

func InitHandler(manager interfaces.IManager) *Handler {
	newHandler := &Handler{
		manager:        manager,
		services:       controllers.InitServicesController(manager),
		serviceTariffs: controllers.InitServiceTariffsController(manager),
		user:           controllers.InitUserController(manager),
	}

	newHandler.InitRoutes()
	return newHandler
}

func (h *Handler) InitRoutes() {
	h.Engine().GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	// here we can add special endpoints
	// based on the Environment
	h.AddRoutesV1()
}

func (h *Handler) Engine() *gin.Engine {
	h.engineRunner.Do(func() {
		h.engine = gin2.InitEngine(h.manager.Cfg().App.Environment)
	})
	return h.engine
}
