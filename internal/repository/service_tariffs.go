package repository

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"github.com/doxanocap/pkg/errs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		return errs.Wrap("repository.serviceTariffs.Create", err)
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
		return nil, errs.Wrap("repository.serviceTariffs.GetByID: ", err)
	}

	if serviceTariff.ID == 0 {
		return nil, nil
	}

	return serviceTariff, nil
}

func (repo *ServiceTariffsRepository) GetAllServices(ctx context.Context) ([]model.ServiceInfo, error) {
	var serviceInfo []model.ServiceInfo
	err := repo.db.WithContext(ctx).
		Select("st.id, s.alias, s.name_key, s.limit, " +
			"s.description_key, st.price, st.auto_pay, st.is_active").
		Table("service_tariffs st").
		Joins("left join service s on st.service_id = s.id").
		Scan(&serviceInfo).
		Error

	if err != nil {
		return nil, errs.Wrap("repository.serviceTariffs.GetByID: ", err)
	}

	if len(serviceInfo) == 0 {
		return nil, nil
	}

	return serviceInfo, nil
}

func (repo *ServiceTariffsRepository) UpdateByID(ctx context.Context, tariff *model.ServiceTariff, tariffID int64) error {
	err := repo.db.
		WithContext(ctx).
		Table("service_tariffs").
		Where("id = ?", tariffID).
		Updates(&tariff).
		Error

	if err != nil {
		return errs.Wrap("repository.serviceTariffs.UpdateByID", err)
	}
	return nil
}

func (repo *ServiceTariffsRepository) DeleteByID(ctx context.Context, tariffID int64) (*model.ServiceTariff, error) {
	var tariff *model.ServiceTariff
	err := repo.db.
		WithContext(ctx).
		Table("service_tariffs").
		Clauses(clause.Returning{}).
		Where("id = ?", tariffID).
		Delete(&tariff).
		Error

	if err != nil {
		return nil, errs.Wrap("repository.serviceTariffs.DeleteByID", err)
	}
	return tariff, nil
}
