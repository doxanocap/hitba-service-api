package interfaces

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
)

type IService interface {
	Services() IServicesService
	User() IUserService
}

// ToDo: rename Services
type IServicesService interface {
	GetAll(ctx context.Context) []model.Service
}

type IUserService interface {
}
