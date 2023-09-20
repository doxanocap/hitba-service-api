package manager

import (
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/repository"
	"github.com/doxanocap/hitba-service-api/pkg/logger"
	"gorm.io/gorm"
	"sync"
)

type RepositoryManager struct {
	db *gorm.DB

	services       interfaces.IServicesRepository
	servicesRunner sync.Once

	serviceTariffs       interfaces.IServiceTariffsRepository
	serviceTariffsRunner sync.Once

	purchasedServices       interfaces.IPurchasedServicesRepository
	purchasedServicesRunner sync.Once
}

func InitRepositoryManager(db *gorm.DB) *RepositoryManager {
	return &RepositoryManager{
		db: db,
	}
}

func (rm *RepositoryManager) Services() interfaces.IServicesRepository {
	rm.servicesRunner.Do(func() {
		rm.services = repository.InitServicesRepository(rm.db, logger.Log.Named("[REPOSITORY][SERVICES]"))
	})
	return rm.services
}

func (rm *RepositoryManager) ServiceTariffs() interfaces.IServiceTariffsRepository {
	rm.serviceTariffsRunner.Do(func() {
		rm.serviceTariffs = repository.InitServiceTariffsRepository(rm.db, logger.Log.Named("[REPOSITORY][SERVICE_TARIFFS]"))
	})
	return rm.serviceTariffs
}

func (rm *RepositoryManager) PurchasedServices() interfaces.IPurchasedServicesRepository {
	rm.purchasedServicesRunner.Do(func() {
		rm.purchasedServices = repository.InitPurchasedServicesRepository(rm.db, logger.Log.Named("[REPOSITORY][PURCHASED_SERVICES]"))
	})
	return rm.purchasedServices
}
