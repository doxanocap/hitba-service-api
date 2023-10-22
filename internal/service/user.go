package service

import (
	"context"
	"github.com/doxanocap/hitba-service-api/internal/manager/interfaces"
	"github.com/doxanocap/hitba-service-api/internal/model"
	"time"
)

type UserService struct {
	manager interfaces.IManager
}

func InitUserService(manager interfaces.IManager) *UserService {
	return &UserService{
		manager: manager,
	}
}

func (s *UserService) PurchaseService(ctx context.Context, purchase *model.Purchase, userID int64) error {
	tariff, err := s.manager.Repository().ServiceTariffs().GetByID(ctx, purchase.TariffID)
	if err != nil {
		return err
	}

	var (
		stamp = time.Now()
	)

	err = s.manager.Processor().BillingAPI().Pay(ctx, &model.PaymentRequest{
		UserID:    userID,
		TariffID:  tariff.ID,
		Price:     tariff.Price,
		CreatedAt: stamp,
	})
	if err != nil {
		return err
	}

	tariff.Limit -= 1
	err = s.manager.Repository().PurchasedServices().Create(ctx, &model.PurchasedService{
		UserID:         userID,
		TariffID:       tariff.ID,
		RemainingLimit: tariff.Limit,
		CreatedAt:      stamp,
	})

	err = s.manager.Repository().ServiceTariffs().UpdateByID(ctx, tariff, tariff.ID)
	if err != nil {
		return err
	}
	return nil
}
