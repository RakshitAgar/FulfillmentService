package controller

//import (
//	"context"
//	"errors"
//	pb "fullfilmentService/proto"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//// MockFulfillmentService is a mock implementation of the FulfillmentService
//type MockFulfillmentService struct {
//	mock.Mock
//}
//
//func (m *MockFulfillmentService) AssignDelivery(ctx context.Context, req *pb.AssignDeliveryPartnerRequest) (*pb.AssignDeliveryResponse, error) {
//	args := m.Called(ctx, req)
//	return args.Get(0).(*pb.AssignDeliveryResponse), args.Error(1)
//}
//
//func TestProcessOrder(t *testing.T) {
//	mockService := new(MockFulfillmentService)
//	controller := NewFulfillmentController(mockService)
//
//	req := &pb.AssignDeliveryPartnerRequest{
//		OrderId: "order-id",
//	}
//
//	deliveryResponse := &pb.AssignDeliveryResponse{
//		DeliveryId: "delivery-id",
//		OrderId:    "order-id",
//		CustomerId: "customer-id",
//		Status:     "assigned",
//	}
//
//	mockService.On("AssignDelivery", mock.Anything, req).Return(deliveryResponse, nil)
//
//	resp, err := controller.ProcessOrder(context.Background(), req)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, resp)
//	assert.Equal(t, "delivery-id", resp.DeliveryId)
//	assert.Equal(t, "order-id", resp.OrderId)
//	assert.Equal(t, "customer-id", resp.CustomerId)
//	assert.Equal(t, "assigned", resp.Status)
//	mockService.AssertExpectations(t)
//}
//
//func TestProcessOrder_AssignDeliveryError(t *testing.T) {
//	mockService := new(MockFulfillmentService)
//	controller := NewFulfillmentController(mockService)
//
//	req := &pb.AssignDeliveryPartnerRequest{
//		OrderId: "order-id",
//	}
//
//	mockService.On("AssignDelivery", mock.Anything, req).Return(nil, errors.New("service error"))
//
//	resp, err := controller.ProcessOrder(context.Background(), req)
//
//	assert.Error(t, err)
//	assert.Nil(t, resp)
//	mockService.AssertExpectations(t)
//}
