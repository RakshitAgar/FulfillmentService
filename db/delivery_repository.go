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

func (r *DeliveryRepository) GetDeliveryByID(ctx context.Context, deliveryID int64) (*model.Delivery, error) {
	var delivery model.Delivery
	if err := r.db.WithContext(ctx).First(&delivery, deliveryID).Error; err != nil {
		return nil, err
	}
	return &delivery, nil
}

func (r *DeliveryRepository) UpdateDelivery(ctx context.Context, delivery *model.Delivery) error {
	return r.db.WithContext(ctx).Save(delivery).Error
}
