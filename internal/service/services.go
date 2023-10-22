package service

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/errs"
)

type ServicesService struct {
	manager interfaces.IManager
}

func InitServicesService(manager interfaces.IManager) *ServicesService {
	return &ServicesService{
		manager: manager,
	}
}

func (s *ServicesService) Create(ctx context.Context, service model.Service) error {
	result, err := s.manager.Repository().Services().GetByName(ctx, service.NameKey)
	if err != nil {
		return errs.Wrap("check if already exists", err)
	}

	if result != nil {
		return model.ErrSuchServiceAlreadyExist
	}

	err = s.manager.Repository().Services().Create(ctx, service)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServicesService) CreateTariff(ctx context.Context, tariff model.ServiceTariff) error {
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

func (s *ServicesService) GetAllServices(ctx context.Context) ([]model.ServiceInfo, error) {
	return s.manager.Repository().ServiceTariffs().GetAllServices(ctx)
}

func (s *ServicesService) GetAll(ctx context.Context) ([]model.Service, error) {
	return s.manager.Repository().Services().GetAll(ctx)
}
