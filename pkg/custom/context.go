package custom

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/consts"
	"sync"
)

const key = "context-holder-key"

func ContextHolderKey() string {
	return key
}

func SetAttribute(ctx context.Context, attribute string, value interface{}) {
	if contextHolder, ok := ctx.Value(key).(*sync.Map); ok {
		contextHolder.Store(attribute, value)
	}
}

func GetAttributeInt64(ctx context.Context, attribute string) int64 {
	value := getAttribute(ctx, attribute)
	if value != nil {
		if i, ok := value.(int64); ok {
			return i
		}
	}
	return consts.NilInt
}

func getAttribute(ctx context.Context, attribute string) interface{} {
	if contextHolder, ok := ctx.Value(key).(*sync.Map); ok {
		value, ok := contextHolder.Load(attribute)
		if ok {
			return value
		}
	}
	return nil
}
