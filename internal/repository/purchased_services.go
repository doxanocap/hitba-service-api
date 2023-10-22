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

func (repo *PurchasedServicesRepository) Create(ctx context.Context, purchase *model.PurchasedService) error {
	err := repo.db.WithContext(ctx).
		Create(&purchase).
		Error

	if err != nil {
		return errs.Wrap("repository.purchasedServices.Create", err)
	}

	return nil
}

func (repo *PurchasedServicesRepository) UpdateByID(ctx context.Context, ps *model.PurchasedService, ID int64) error {
	err := repo.db.
		WithContext(ctx).
		Table("purchased_services").
		Where("id = ?", ID).
		Updates(&ps).
		Error

	if err != nil {
		return errs.Wrap("repository.purchasedServices.UpdateByID", err)
	}
	return nil
}

func (repo *PurchasedServicesRepository) DeleteByID(ctx context.Context, ID int64) (*model.PurchasedService, error) {
	var service *model.PurchasedService
	err := repo.db.
		WithContext(ctx).
		Table("purchased_services").
		Where("id = ?", ID).
		Delete(&service).
		Error

	if err != nil {
		return nil, errs.Wrap("repository.purchasedServices.DeleteByID", err)
	}
	return service, nil
}
