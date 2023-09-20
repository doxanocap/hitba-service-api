package processor

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces/processor/rest"
)

type IRESTProcessor interface {
	Handler() rest.IHandlerManager
	Middlewares() rest.IMiddlewareManager
}
