package interfaces

import "github.com/doxanocap/hitba-service-api/internal/manager/interfaces/processor"

type IProcessor interface {
	REST() processor.IRESTProcessor
}
