package model

import "time"

type Service struct {
	ID             int64     `json:"id" db:"id"`
	Alias          string    `json:"alias" db:"alias"`
	NameKey        string    `json:"name_key" db:"name_key"`
	PricePerUnit   int       `json:"price_per_unit" db:"price_per_unit"`
	DescriptionKey string    `json:"description_key" db:"description_key"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type ServiceTariff struct {
	ID             int64     `json:"id" db:"id"`
	ServiceID      int64     `json:"service_id" db:"service_id"`
	Limit          int       `json:"limit" db:"limit"`
	LimitationType string    `json:"limitation_type" db:"limitation_type"`
	Price          int       `json:"price" db:"price"`
	AutoPay        bool      `json:"auto_pay" db:"auto_pay"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type PurchasedService struct {
	ID             int64     `json:"id" db:"id"`
	UserID         int64     `json:"user_id" db:"user_id"`
	TariffID       int       `json:"tariff_id" db:"tariff_id"`
	RemainingLimit int       `json:"remaining_limit" db:"remaining_limit"`
	ExpireAt       time.Time `json:"expire_at" db:"expire_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
