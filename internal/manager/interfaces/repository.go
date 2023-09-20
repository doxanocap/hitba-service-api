package interfaces

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
)

type IRepository interface {
	Services() IServicesRepository
}

type IServicesRepository interface {
	GetAll(ctx context.Context) []model.Service
}

type IServiceTariffsRepository interface {
}

type IPurchasedServicesRepository interface {
}
