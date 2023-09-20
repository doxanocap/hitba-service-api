package service

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
)

type UserService struct {
	manager interfaces.IManager
}

func InitUserService(manager interfaces.IManager) *UserService {
	return &UserService{
		manager: manager,
	}
}
