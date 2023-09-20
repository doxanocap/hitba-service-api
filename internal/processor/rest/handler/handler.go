package handler

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest/utils"
	"github.com/gin-gonic/gin"
	"sync"
)

type Handler struct {
	services *ServicesController
	user     *UserController

	engine       *gin.Engine
	engineRunner sync.Once
	manager      interfaces.IManager
}

func InitHandler(manager interfaces.IManager) *Handler {
	newHandler := &Handler{
		manager:  manager,
		services: InitServicesController(manager),
		user:     InitUserController(manager),
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
		h.engine = utils.InitEngine(h.manager.Cfg().App.Environment)
	})
	return h.engine
}
