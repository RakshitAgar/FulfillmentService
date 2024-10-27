package model

import (
	"testing"
)

func TestNewDelivery_ValidCases(t *testing.T) {
	tests := []struct {
		orderID    int64
		customerID string
		status     string
	}{
		{1, "customer1", "assigned"},
		{2, "customer2", "delivered"},
	}

	for _, tt := range tests {
		delivery, err := NewDelivery(tt.orderID, tt.customerID, tt.status)
		if err != nil {
			t.Errorf("did not expect error but got %v", err)
		}
		if delivery.OrderID != tt.orderID {
			t.Errorf("expected orderID %d but got %d", tt.orderID, delivery.OrderID)
		}
		if delivery.CustomerID != tt.customerID {
			t.Errorf("expected customerID %s but got %s", tt.customerID, delivery.CustomerID)
		}
		if delivery.Status != tt.status {
			t.Errorf("expected status %s but got %s", tt.status, delivery.Status)
		}
	}
}

func TestNewDelivery_InvalidCases(t *testing.T) {
	tests := []struct {
		orderID    int64
		customerID string
		status     string
	}{
		{0, "customer1", "assigned"},
		{1, "", "assigned"},
		{2, "customer2", ""},
	}

	for _, tt := range tests {
		_, err := NewDelivery(tt.orderID, tt.customerID, tt.status)
		if err == nil {
			t.Errorf("expected error but got none")
		}
	}
}
