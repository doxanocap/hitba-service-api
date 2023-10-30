package handler

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest/controllers"
	"github.com/doxanocap/hitba-service-api/pkg/custom"
	"github.com/gin-gonic/gin"
	"sync"
)

type Handler struct {
	services *controllers.ServicesController
	user     *controllers.UserController

	manager interfaces.IManager
	config  *model.Config

	engine       *gin.Engine
	engineRunner sync.Once
}

func InitHandler(manager interfaces.IManager, config *model.Config) *Handler {
	newHandler := &Handler{
		manager:  manager,
		config:   config,
		services: controllers.InitServicesController(manager, config),
		user:     controllers.InitUserController(manager, config),
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
		h.engine = custom.InitEngine(h.config.App.Environment)
	})
	return h.engine
}
