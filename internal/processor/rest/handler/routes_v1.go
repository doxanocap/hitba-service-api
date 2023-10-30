package handler

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine().Group("/v1")

	// list of all paid services
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	services := v1.Group("services")
	{
		services.POST("", h.services.Create)
		services.GET("", h.services.GetAll)
	}

	serviceTariffs := v1.Group("tariffs")
	{
		serviceTariffs.POST("", h.services.CreateTariff)
		serviceTariffs.PUT("/:id", h.services.UpdateTariff)
	}

	users := v1.Group("users")
	{
		users.POST("/:user_id/buy", h.user.PurchaseServiceByID)
	}
}
