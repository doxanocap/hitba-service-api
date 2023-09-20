package manager

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/service"
	"sync"
)

type ServiceManager struct {
	manager interfaces.IManager

	services       interfaces.IServicesService
	servicesRunner sync.Once

	user       interfaces.IUserService
	userRunner sync.Once
}

func InitServiceManager(manager interfaces.IManager) *ServiceManager {
	return &ServiceManager{
		manager: manager,
	}
}

func (s *ServiceManager) Services() interfaces.IServicesService {
	s.servicesRunner.Do(func() {
		s.services = service.InitServicesService(s.manager)
	})
	return s.services
}

func (s *ServiceManager) User() interfaces.IUserService {
	s.userRunner.Do(func() {
		s.user = service.InitUserService(s.manager)
	})
	return s.user
}
