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
		return errs.Wrap("repo_services.Create", err)
	}
	return nil
}

func (repo *ServicesRepository) GetAll(ctx context.Context) ([]model.Service, error) {
	var services []model.Service

	err := repo.db.WithContext(ctx).
		Raw("SELECT * FROM services").
		Scan(&services).
		Error
	if err != nil {
		return nil, errs.Wrap("repo_services.GetAll", err)
	}
	return services, nil
}

func (repo *ServicesRepository) GetByName(ctx context.Context, nameKey string) (*model.Service, error) {
	service := &model.Service{}
	err := repo.db.WithContext(ctx).
		Raw("SELECT * FROM services WHERE name_key = ? Limit 1", nameKey).
		Scan(service).
		Error

	if err != nil {
		return nil, errs.Wrap("repo_services.GetByName", err)
	}

	if service.ID == 0 {
		return nil, nil
	}

	return service, nil
}

func (repo *ServicesRepository) GetByID(ctx context.Context, ID int64) (*model.Service, error) {
	service := &model.Service{}
	err := repo.db.WithContext(ctx).
		Raw("SELECT * FROM services WHERE id = ?", ID).
		Scan(service).
		Error

	if err != nil {
		return nil, errs.Wrap("repo_services.GetByID: ", err)
	}

	if service.ID == 0 {
		return nil, nil
	}

	return service, nil
}
