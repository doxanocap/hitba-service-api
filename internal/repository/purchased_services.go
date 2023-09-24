package repository

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/errs"
	"gorm.io/gorm"
)

type PurchasedServicesRepository struct {
	db *gorm.DB
}

func InitPurchasedServicesRepository(db *gorm.DB) *PurchasedServicesRepository {
	return &PurchasedServicesRepository{
		db: db,
	}
}

func (repo *PurchasedServicesRepository) Create(ctx context.Context, purchasedService model.PurchasedService) error {
	err := repo.db.WithContext(ctx).
		Create(&purchasedService).
		Error

	if err != nil {
		return errs.Wrap("repo_purchasedServices.Create", err)
	}

	return nil
}
