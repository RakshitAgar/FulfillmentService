package controller

import (
	"context"
	"fullfilmentService/internal/service"
	pb "fullfilmentService/proto"
)

type FulfillmentController struct {
	pb.UnimplementedFulfillmentServiceServer
	service *service.FulfillmentService
}

func NewFulfillmentController(service *service.FulfillmentService) *FulfillmentController {
	return &FulfillmentController{service: service}
}

/*
1. Call the AssignDelivery method from the service
2. Check for the Error and return it if it exists
3. Create and return a response with the delivery ID, order ID, customer ID, and status
*/

func (h *FulfillmentController) ProcessOrder(ctx context.Context, req *pb.AssignDeliveryPartnerRequest) (*pb.AssignDeliveryResponse, error) {
	delivery, err := h.service.AssignDelivery(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.AssignDeliveryResponse{
		DeliveryId: delivery.DeliveryId,
		OrderId:    delivery.OrderId,
		CustomerId: delivery.CustomerId,
		Status:     delivery.Status,
	}, nil
}
