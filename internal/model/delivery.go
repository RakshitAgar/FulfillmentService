package model

import "errors"

type Delivery struct {
	DeliveryID      int64 `gorm:"primaryKey;autoIncrement"`
	OrderID         int64
	CustomerID      int64
	Status          string
	DeliveryAddress string
}

func NewDelivery(orderID int64, customerID int64, status string, deliveryAddress string) (*Delivery, error) {
	if orderID == 0 {
		return nil, errors.New("orderID cannot be zero")
	}
	if customerID == 0 {
		return nil, errors.New("customerID cannot be empty")
	}
	if status == "" || deliveryAddress == "" {
		return nil, errors.New("status or deliveryAddress cannot be empty")
	}
	return &Delivery{
		OrderID:         orderID,
		CustomerID:      customerID,
		Status:          status,
		DeliveryAddress: deliveryAddress,
	}, nil
}
