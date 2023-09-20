package rest

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces/processor/rest"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest/handler"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest/middlewares"
	"sync"
)

type REST struct {
	manager interfaces.IManager

	handler       rest.IHandlerManager
	handlerRunner sync.Once

	middlewares       rest.IMiddlewareManager
	middlewaresRunner sync.Once
}

func Init(manager interfaces.IManager) *REST {
	return &REST{
		manager: manager,
	}
}

func (r *REST) Handler() rest.IHandlerManager {
	r.handlerRunner.Do(func() {
		r.handler = handler.InitHandler(r.manager)
	})
	return r.handler
}

func (r *REST) Middlewares() rest.IMiddlewareManager {
	r.middlewaresRunner.Do(func() {
		r.middlewares = middlewares.InitMiddlewares(r.manager)
	})
	return r.middlewares
}
