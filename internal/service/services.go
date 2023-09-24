package service

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/consts"
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
	result, err := s.manager.Repository().Services().FindByName(ctx, service.NameKey)
	if err != nil {
		return errs.Wrap("idempotency check", err)
	}

	if result != nil {
		return consts.ErrSuchServiceAlreadyExist
	}

	err = s.manager.Repository().Services().Create(ctx, service)
	if err != nil {
		return errs.Wrap("create new service", err)
	}
	return nil
}

func (s *ServicesService) GetAll(ctx context.Context) ([]model.Service, error) {
	return s.manager.Repository().Services().GetAll(ctx)
}
