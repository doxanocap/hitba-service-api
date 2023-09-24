package interfaces

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
)

type IRepository interface {
	Services() IServicesRepository
}

type IServicesRepository interface {
	Create(ctx context.Context, service model.Service) error
	FindByName(ctx context.Context, nameKey string) (*model.Service, error)
	GetAll(ctx context.Context) ([]model.Service, error)
}

type IServiceTariffsRepository interface {
}

type IPurchasedServicesRepository interface {
}
