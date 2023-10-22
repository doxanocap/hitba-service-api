package manager

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/hitba-service-api/internal/processor/api"
	"github.com/doxanocap/hitba-service-api/internal/processor/rest"
	"sync"
)

type ProcessorManager struct {
	manager interfaces.IManager
	config  *model.Config

	restProcessor       interfaces.IRESTProcessor
	restProcessorRunner sync.Once

	billingAPIProcessor       interfaces.IBillingAPIProcessor
	billingAPIProcessorRunner sync.Once
}

func InitProcessor(manager interfaces.IManager, config *model.Config) *ProcessorManager {
	return &ProcessorManager{
		manager: manager,
		config:  config,
	}
}

func (p *ProcessorManager) REST() interfaces.IRESTProcessor {
	p.restProcessorRunner.Do(func() {
		p.restProcessor = rest.Init(p.manager)
	})
	return p.restProcessor
}

func (p *ProcessorManager) BillingAPI() interfaces.IBillingAPIProcessor {
	p.billingAPIProcessorRunner.Do(func() {
		p.billingAPIProcessor = api.NewBillingAPI(p.manager, p.config)
	})
	return p.billingAPIProcessor
}
