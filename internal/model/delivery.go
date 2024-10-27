package model

import "errors"

type Delivery struct {
	DeliveryID int64 `gorm:"primaryKey;autoIncrement"`
	OrderID    int64
	CustomerID string
	Status     string
}

func NewDelivery(orderID int64, customerID, status string) (*Delivery, error) {
	if orderID == 0 {
		return nil, errors.New("orderID cannot be zero")
	}
	if customerID == "" {
		return nil, errors.New("customerID cannot be empty")
	}
	if status == "" {
		return nil, errors.New("status cannot be empty")
	}
	return &Delivery{
		OrderID:    orderID,
		CustomerID: customerID,
		Status:     status,
	}, nil
}
