package repository

import (
	"context"
	"errors"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/errs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServicesRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func InitServicesRepository(db *gorm.DB, log *zap.Logger) *ServicesRepository {
	return &ServicesRepository{
		db:  db,
		log: log,
	}
}

func (repo *ServicesRepository) Create(ctx context.Context, service model.Service) error {
	err := repo.db.WithContext(ctx).
		Model(&model.Service{}).
		Create(service).Error
	return errs.Wrap("servicesRepo.Create", err)
}

func (repo *ServicesRepository) GetAll(ctx context.Context) ([]model.Service, error) {
	var services []model.Service

	err := repo.db.WithContext(ctx).
		Model(&model.Service{}).
		Select("*").
		Scan(&services).
		Error
	return services, err
}

func (repo *ServicesRepository) FindByName(ctx context.Context, nameKey string) (*model.Service, error) {
	var service model.Service
	err := repo.db.WithContext(ctx).
		Model(&model.Service{}).
		Where("name_key = ?", nameKey).
		First(&service).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, errs.Wrap("servicesRepo.FindByName: ", err)
	}

	return &service, nil
}
