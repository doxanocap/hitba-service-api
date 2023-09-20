package manager

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	IProcessor "github.com/doxanocap/hitba-service-api/internal/manager/interfaces/processor"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest"
	"sync"
)

type ProcessorManager struct {
	manager interfaces.IManager

	restProcessor       IProcessor.IRESTProcessor
	restProcessorRunner sync.Once
}

func InitProcessor(manager interfaces.IManager) *ProcessorManager {
	return &ProcessorManager{
		manager: manager,
	}
}

func (p *ProcessorManager) REST() IProcessor.IRESTProcessor {
	p.restProcessorRunner.Do(func() {
		p.restProcessor = rest.Init(p.manager)
	})
	return p.restProcessor
}
