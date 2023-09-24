package service

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/errs"
)

type ServicesTariffsService struct {
	manager interfaces.IManager
}

func InitServicesTariffsService(manager interfaces.IManager) *ServicesTariffsService {
	return &ServicesTariffsService{
		manager: manager,
	}
}

func (s *ServicesTariffsService) Create(ctx context.Context, tariff model.ServiceTariff) error {
	result, err := s.manager.Repository().Services().GetByID(ctx, tariff.ServiceID)
	if err != nil {
		return errs.Wrap("check service with such id exist", err)
	}

	if result == nil {
		return model.ErrServiceIdNotFound
	}

	err = s.manager.Repository().ServiceTariffs().Create(ctx, tariff)
	if err != nil {
		return err
	}
	return nil
}
