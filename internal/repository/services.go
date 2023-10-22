package repository

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/errs"
	"gorm.io/gorm"
	"time"
)

type ServicesRepository struct {
	db *gorm.DB
}

func InitServicesRepository(db *gorm.DB) *ServicesRepository {
	return &ServicesRepository{
		db: db,
	}
}

func (repo *ServicesRepository) Create(ctx context.Context, service model.Service) error {
	service.UpdatedAt = time.Now()
	err := repo.db.WithContext(ctx).
		Create(&service).
		Error

	if err != nil {
		return errs.Wrap("repository.services.Create", err)
	}
	return nil
}

func (repo *ServicesRepository) GetAll(ctx context.Context) ([]model.Service, error) {
	var services []model.Service
	err := repo.db.WithContext(ctx).
		Select("*").
		Table("service").
		Scan(&services).
		Error

	if err != nil {
		return nil, errs.Wrap("repository.services.GetAll", err)
	}
	return services, nil
}

func (repo *ServicesRepository) GetByName(ctx context.Context, nameKey string) (*model.Service, error) {
	var service *model.Service
	err := repo.db.WithContext(ctx).
		Select("*").
		Table("service").
		Where("name_key = ?", nameKey).
		Limit(1).
		Scan(&service).
		Error

	if err != nil {
		return nil, errs.Wrap("repository.services.GetByName", err)
	}
	if service.ID == 0 {
		return nil, nil
	}
	return service, nil
}

func (repo *ServicesRepository) GetByID(ctx context.Context, ID int64) (*model.Service, error) {
	var service *model.Service
	err := repo.db.WithContext(ctx).
		Select("*").
		Table("service").
		Where("id = ?", ID).
		Limit(1).
		Scan(&service).
		Error

	if err != nil {
		return nil, errs.Wrap("repository.services.GetByID: ", err)
	}
	if service.ID == 0 {
		return nil, nil
	}
	return service, nil
}

func (repo *ServicesRepository) UpdateByID(ctx context.Context, service *model.Service, ID int64) error {
	err := repo.db.
		WithContext(ctx).
		Table("services").
		Where("id = ?", ID).
		Updates(&service).
		Error

	if err != nil {
		return errs.Wrap("repository.services.UpdateByID", err)
	}
	return nil
}

func (repo *ServicesRepository) DeleteByID(ctx context.Context, ID int64) (*model.Service, error) {
	var service *model.Service
	err := repo.db.
		WithContext(ctx).
		Table("services").
		Where("id = ?", ID).
		Delete(&service).
		Error

	if err != nil {
		return nil, errs.Wrap("repository.services.DeleteByID", err)
	}
	return service, nil
}
