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

// AssignDelivery should match the proto service method name
func (h *FulfillmentController) AssignDelivery(ctx context.Context, req *pb.AssignDeliveryPartnerRequest) (*pb.AssignDeliveryResponse, error) {
	delivery, err := h.service.AssignDeliveryAgent(ctx, req)
	if err != nil {
		return nil, err
	}

	return delivery, nil
}

// UpdateDeliveryStatus should also match the proto service method name
func (h *FulfillmentController) UpdateDeliveryStatus(ctx context.Context, req *pb.UpdateDeliveryStatusRequest) (*pb.UpdateDeliveryStatusResponse, error) {
	response, err := h.service.StatusUpdate(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
