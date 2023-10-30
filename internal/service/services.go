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

func (s *ServicesService) Create(ctx context.Context, service *model.Service) error {
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

func (s *ServicesService) CreateTariff(ctx context.Context, tariff *model.ServiceTariff) error {
	result, err := s.manager.Repository().Services().GetByID(ctx, tariff.ServiceID)
	if err != nil {
		return err
	}

	if result == nil {
		return model.ErrServiceIDNotFound
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

func (s *ServicesService) UpdateTariffByID(ctx context.Context, tariff *model.ServiceTariff, ID int64) error {
	prevTariff, err := s.manager.Repository().ServiceTariffs().GetByID(ctx, ID)
	if err != nil {
		return err
	}
	if prevTariff == nil {
		return model.ErrTariffIDNotFound
	}

	result, err := s.manager.Repository().Services().GetByID(ctx, ID)
	if err != nil {
		return err
	}

	if result == nil {
		return model.ErrServiceIDNotFound
	}

	err = s.manager.Repository().ServiceTariffs().UpdateByID(ctx, tariff, ID)
	if err != nil {
		return err
	}
	return nil
}
