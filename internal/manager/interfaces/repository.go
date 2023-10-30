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
	Create(ctx context.Context, service *model.Service) error
	GetAll(ctx context.Context) ([]model.Service, error)
	GetByID(ctx context.Context, ID int64) (*model.Service, error)
	GetByName(ctx context.Context, nameKey string) (*model.Service, error)
	UpdateByID(ctx context.Context, service *model.Service, ID int64) error
	DeleteByID(ctx context.Context, ID int64) (*model.Service, error)
}

type IServiceTariffsRepository interface {
	Create(ctx context.Context, serviceTariff *model.ServiceTariff) error
	GetByID(ctx context.Context, ID int64) (*model.ServiceTariff, error)
	GetAllServices(ctx context.Context) ([]model.ServiceInfo, error)
	UpdateByID(ctx context.Context, tariff *model.ServiceTariff, tariffID int64) error
	DeleteByID(ctx context.Context, tariffID int64) (*model.ServiceTariff, error)
}

type IPurchasedServicesRepository interface {
	Create(ctx context.Context, purchase *model.PurchasedService) error
}
