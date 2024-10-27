package service

import (
	"context"
	"fullfilmentService/db"
	"fullfilmentService/internal/model"
	pb "fullfilmentService/proto"
	"log"
)

type FulfillmentService struct {
	pb.UnimplementedFulfillmentServiceServer
	repo *db.DeliveryRepository
}

func NewFulfillmentService(repo *db.DeliveryRepository) *FulfillmentService {
	return &FulfillmentService{repo: repo}
}

/*
1. Generate a delivery ID
2. Set the status to "assigned"
3. Create a new Delivery object
4. Save the delivery to the database
5. Return a response with the delivery ID, order ID, customer ID, and status
*/

func (s *FulfillmentService) AssignDelivery(ctx context.Context, req *pb.AssignDeliveryPartnerRequest) (*pb.AssignDeliveryResponse, error) {
	status := "assigned"

	// have to add call for getting the order detail from the order service with the order ID Only
	// and then get the customer ID from the order detail

	delivery, errInvalidCredentials := model.NewDelivery(req.OrderId, "", status)

	if errInvalidCredentials != nil {
		return nil, errInvalidCredentials
	}

	err := s.repo.SaveDelivery(ctx, delivery)
	if err != nil {
		return nil, err
	}

	resp := &pb.AssignDeliveryResponse{
		DeliveryId: delivery.DeliveryID,
		OrderId:    req.OrderId,
		CustomerId: "customer-id",
		Status:     status,
	}

	log.Printf("Assigned delivery: %v", resp)
	return resp, nil
}
