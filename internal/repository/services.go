package repository

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
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
	return repo.db.Model(&model.Service{}).Create(service).Error
}

func (repo *ServicesRepository) GetAll(ctx context.Context) []model.Service {

	return make([]model.Service, 0)
}
