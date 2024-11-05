// internal/client/orderClient_test.go
package client

import (
	_ "bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrderDetail(t *testing.T) {
	orderID := int64(1)
	expectedOrderDetail := &OrderDetail{
		ID:              orderID,
		RestaurantID:    1,
		CustomerID:      1,
		TotalPrice:      100.0,
		Status:          "PENDING",
		OrderItems:      []OrderItem{},
		DeliveryAddress: "123 Test St",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/orders/1" {
			t.Fatalf("expected path to be /orders/1, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedOrderDetail)
	}))
	defer server.Close()

	orderDetail, err := GetOrderDetail(orderID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if orderDetail.ID != expectedOrderDetail.ID {
		t.Errorf("expected order ID %d, got %d", expectedOrderDetail.ID, orderDetail.ID)
	}
}

func TestUpdateOrderStatus(t *testing.T) {
	orderID := int64(1)
	status := "DELIVERED"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/orders/1/status" {
			t.Fatalf("expected path to be /orders/1/status, got %s", r.URL.Path)
		}
		var statusUpdate map[string]string
		if err := json.NewDecoder(r.Body).Decode(&statusUpdate); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if statusUpdate["status"] != status {
			t.Errorf("expected status %s, got %s", status, statusUpdate["status"])
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	err := UpdateOrderStatus(orderID, status)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
