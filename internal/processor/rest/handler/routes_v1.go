package handler

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine().Group("/v1")

	// list of all paid services
	services := v1.Group("services")
	{
		services.GET("", h.services.GetAll)
	}

	users := v1.Group("users")
	{
		users.POST("/:user_id/buy")
	}
}
