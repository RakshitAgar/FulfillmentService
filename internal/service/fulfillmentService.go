// internal/service/fulfillmentService.go
package service

import (
	"context"
	"fmt"
	"fullfilmentService/db"
	"fullfilmentService/internal/client"
	"fullfilmentService/internal/model"
	pb "fullfilmentService/proto"
)

type FulfillmentService struct {
	pb.UnimplementedFulfillmentServiceServer
	repo *db.DeliveryRepository
}

func NewFulfillmentService(repo *db.DeliveryRepository) *FulfillmentService {
	return &FulfillmentService{repo: repo}
}

/*
AssignDelivery assigns a delivery partner to an order.
It calls the order service to get the order details and then saves the delivery details in the database.
*/

func (s *FulfillmentService) AssignDeliveryAgent(ctx context.Context, req *pb.AssignDeliveryPartnerRequest) (*pb.AssignDeliveryResponse, error) {
	status := "assigned"

	// Call the order service to get order details
	orderDetail, err := client.GetOrderDetail(req.OrderId)
	if err != nil {
		return nil, err
	}

	delivery, errInvalidCredentials := model.NewDelivery(req.OrderId, orderDetail.CustomerID, status, orderDetail.DeliveryAddress)
	if errInvalidCredentials != nil {
		return nil, errInvalidCredentials
	}

	err = s.repo.SaveDelivery(ctx, delivery)
	if err != nil {
		return nil, err
	}

	resp := &pb.AssignDeliveryResponse{
		DeliveryId:       delivery.DeliveryID,
		OrderId:          req.OrderId,
		CustomerId:       orderDetail.CustomerID,
		Delivery_Address: orderDetail.DeliveryAddress,
		Status:           status,
	}

	fmt.Printf("Assigned delivery: %+v\n", resp)
	return resp, nil
}

/*
1. Fetch the delivery details using the delivery ID
2. Update the status of the delivery
3. Save the updated delivery details
4. Return the updated status
*/

func (s *FulfillmentService) StatusUpdate(ctx context.Context, req *pb.UpdateDeliveryStatusRequest) (*pb.UpdateDeliveryStatusResponse, error) {
	// Fetch the delivery details using the delivery ID
	delivery, err := s.repo.GetDeliveryByID(ctx, req.DeliveryId)
	if err != nil {
		return nil, err
	}

	// Update the status of the delivery
	delivery.Status = req.Status

	// should we also update the order status ?
	// Call the order service to get order details

	// Save the updated delivery details
	err = s.repo.UpdateDelivery(ctx, delivery)
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateDeliveryStatusResponse{
		Updated_Status: delivery.Status,
	}

	return resp, nil
}
