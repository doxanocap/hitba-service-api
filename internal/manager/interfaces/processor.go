package interfaces

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces/processor"
	"github.com/doxanocap/hitba-service-api/internal/model"
)

type IProcessor interface {
	REST() IRESTProcessor
	BillingAPI() IBillingAPIProcessor
}

type IRESTProcessor interface {
	Handler() processor.IHandlerManager
	Middlewares() processor.IMiddlewareManager
}

type IBillingAPIProcessor interface {
	Pay(ctx context.Context, payment *model.PaymentRequest) error
}
