package rest

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces/processor"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest/handler"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest/middlewares"
	"sync"
)

type REST struct {
	manager interfaces.IManager
	config  *model.Config

	handler       processor.IHandlerManager
	handlerRunner sync.Once

	middlewares       processor.IMiddlewareManager
	middlewaresRunner sync.Once
}

func Init(manager interfaces.IManager, config *model.Config) *REST {
	return &REST{
		manager: manager,
		config:  config,
	}
}

func (r *REST) Handler() processor.IHandlerManager {
	r.handlerRunner.Do(func() {
		r.handler = handler.InitHandler(r.manager, r.config)
	})
	return r.handler
}

func (r *REST) Middlewares() processor.IMiddlewareManager {
	r.middlewaresRunner.Do(func() {
		r.middlewares = middlewares.InitMiddlewares(r.manager)
	})
	return r.middlewares
}
