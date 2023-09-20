package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PurchasedServicesRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func InitPurchasedServicesRepository(db *gorm.DB, log *zap.Logger) *PurchasedServicesRepository {
	return &PurchasedServicesRepository{
		db:  db,
		log: log,
	}
}
