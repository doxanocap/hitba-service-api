package model

import "time"

type PaymentRequest struct {
	UserID    int64     `json:"user_id" db:"user_id"`
	TariffID  int64     `json:"tariff_id" db:"tariff_id"`
	Price     int       `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
