package service

//
//import (
//	"context"
//	"errors"
//	"fullfilmentService/db"
//	"fullfilmentService/internal/model"
//	pb "fullfilmentService/proto"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//// Move MockDeliveryRepository to db package
//type mockDeliveryRepository struct {
//	mock.Mock
//	*db.DeliveryRepository
//}
//
//func (m *mockDeliveryRepository) SaveDelivery(ctx context.Context, delivery *model.Delivery) error {
//	args := m.Called(ctx, delivery)
//	return args.Error(0)
//}
//
//func TestAssignDelivery(t *testing.T) {
//	mockRepo := &mockDeliveryRepository{}
//	service := NewFulfillmentService(mockRepo)
//
//	req := &pb.AssignDeliveryPartnerRequest{
//		OrderId:    12345,
//		DeliveryId: 0,
//	}
//
//	// Set expectations for SaveDelivery method in the mock
//	mockRepo.On("SaveDelivery", mock.Anything, mock.AnythingOfType("*model.Delivery")).Return(nil)
//
//	// Call the AssignDelivery method
//	resp, err := service.AssignDelivery(context.Background(), req)
//
//	// Validate response and expectations
//	assert.NoError(t, err)
//	assert.NotNil(t, resp)
//	assert.Equal(t, req.OrderId, resp.OrderId)
//	assert.Equal(t, "customer-id", resp.CustomerId)
//	assert.Equal(t, "assigned", resp.Status)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestAssignDelivery_SaveDeliveryError(t *testing.T) {
//	mockRepo := &mockDeliveryRepository{}
//	service := NewFulfillmentService(mockRepo)
//
//	req := &pb.AssignDeliveryPartnerRequest{
//		OrderId:    12345,
//		DeliveryId: 0,
//	}
//
//	// Simulate a failure in the SaveDelivery method
//	mockRepo.On("SaveDelivery", mock.Anything, mock.AnythingOfType("*model.Delivery")).Return(errors.New("db error"))
//
//	// Call the AssignDelivery method
//	resp, err := service.AssignDelivery(context.Background(), req)
//
//	// Validate response and expectations
//	assert.Error(t, err)
//	assert.Nil(t, resp)
//	mockRepo.AssertExpectations(t)
//}
