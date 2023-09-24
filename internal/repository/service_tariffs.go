package repository

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/errs"
	"gorm.io/gorm"
	"time"
)

type ServiceTariffsRepository struct {
	db *gorm.DB
}

func InitServiceTariffsRepository(db *gorm.DB) *ServiceTariffsRepository {
	return &ServiceTariffsRepository{
		db: db,
	}
}

func (repo *ServiceTariffsRepository) Create(ctx context.Context, tariff model.ServiceTariff) error {
	tariff.UpdatedAt = time.Now()

	err := repo.db.WithContext(ctx).
		Create(&tariff).
		Error

	if err != nil {
		return errs.Wrap("repo_serviceTariffs.Create", err)
	}

	return nil
}

func (repo *ServiceTariffsRepository) GetByID(ctx context.Context, ID int64) (*model.ServiceTariff, error) {
	serviceTariff := &model.ServiceTariff{}

	err := repo.db.WithContext(ctx).
		Raw("SELECT * FROM service_tariffs WHERE id = ?", ID).
		Scan(serviceTariff).
		Error

	if err != nil {
		return nil, errs.Wrap("repo_serviceTariffs.GetByID: ", err)
	}

	if serviceTariff.ID == 0 {
		return nil, nil
	}

	return serviceTariff, nil
}
