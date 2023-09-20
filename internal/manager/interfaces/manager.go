package interfaces

import "github.com/doxanocap/hitba-service-api/internal/model"

type IManager interface {
	Processor() IProcessor
	Service() IService
	Repository() IRepository

	Cfg() *model.Config
}
