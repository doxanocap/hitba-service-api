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
	Create(ctx context.Context, service model.Service) error
	CreateTariff(ctx context.Context, tariff model.ServiceTariff) error
	GetAllServices(ctx context.Context) ([]model.ServiceInfo, error)
	GetAll(ctx context.Context) ([]model.Service, error)
}

type IUserService interface {
	PurchaseService(ctx context.Context, purchase *model.Purchase, userID int64) error
}
