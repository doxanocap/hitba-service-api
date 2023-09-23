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
	err := s.manager.Repository().Services().Create(ctx, service)
	if err != nil {
		return errs.Wrap("create service", err)
	}
	return nil
}

func (s *ServicesService) GetAll(ctx context.Context) []model.Service {
	return s.manager.Repository().Services().GetAll(ctx)
}
