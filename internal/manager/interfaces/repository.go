package interfaces

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
)

type IRepository interface {
	Services() IServicesRepository
	ServiceTariffs() IServiceTariffsRepository
	PurchasedServices() IPurchasedServicesRepository
}

type IServicesRepository interface {
	Create(ctx context.Context, service model.Service) error
	GetAll(ctx context.Context) ([]model.Service, error)
	GetByID(ctx context.Context, ID int64) (*model.Service, error)
	GetByName(ctx context.Context, nameKey string) (*model.Service, error)
}

type IServiceTariffsRepository interface {
	Create(ctx context.Context, serviceTariff model.ServiceTariff) error
	GetByID(ctx context.Context, ID int64) (*model.ServiceTariff, error)
}

type IPurchasedServicesRepository interface {
	Create(ctx context.Context, purchasedService model.PurchasedService) error
}
