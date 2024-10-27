package db

import (
	"context"
	"fullfilmentService/internal/model"
	"gorm.io/gorm"
)

type DeliveryRepository struct {
	db *gorm.DB
}

func NewDeliveryRepository(db *gorm.DB) *DeliveryRepository {
	return &DeliveryRepository{db: db}
}

func (r *DeliveryRepository) SaveDelivery(ctx context.Context, delivery *model.Delivery) error {
	return r.db.WithContext(ctx).Create(delivery).Error
}
