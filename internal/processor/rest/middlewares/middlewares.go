package middlewares

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces/processor/rest"
)

type Middlewares struct {
	manager interfaces.IManager
}

func InitMiddlewares(manager interfaces.IManager) rest.IMiddlewareManager {
	return &Middlewares{
		manager: manager,
	}
}
