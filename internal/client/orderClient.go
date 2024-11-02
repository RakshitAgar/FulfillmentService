// internal/client/orderClient.go
package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderItem struct {
	ID           int64   `json:"id"`
	MenuItemID   int64   `json:"menuItemId"`
	MenuItemName string  `json:"menuItemName"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
}

type OrderDetail struct {
	ID              int64       `json:"id"`
	RestaurantID    int64       `json:"restaurantId"`
	CustomerID      int64       `json:"customerId"`
	TotalPrice      float64     `json:"totalPrice"`
	Status          string      `json:"status"`
	OrderItems      []OrderItem `json:"orderItems"`
	DeliveryAddress string      `json:"deliveryAddress"`
}

func GetOrderDetail(orderID int64) (*OrderDetail, error) {
	url := fmt.Sprintf("http://localhost:8083/orders/%d", orderID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call order service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("order service returned status: %s", resp.Status)
	}

	var orderDetail OrderDetail
	if err := json.NewDecoder(resp.Body).Decode(&orderDetail); err != nil {
		return nil, fmt.Errorf("failed to decode order service response: %w", err)
	}

	return &orderDetail, nil
}
