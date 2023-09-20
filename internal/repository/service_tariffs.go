package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceTariffsRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func InitServiceTariffsRepository(db *gorm.DB, log *zap.Logger) *ServiceTariffsRepository {
	return &ServiceTariffsRepository{
		db:  db,
		log: log,
	}
}
