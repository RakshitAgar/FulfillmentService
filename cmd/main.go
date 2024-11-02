package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"fullfilmentService/db"
	"fullfilmentService/internal/controller"
	"fullfilmentService/internal/service"
	pb "fullfilmentService/proto"
	"google.golang.org/grpc"
)

func main() {
	// Initialize the database connection
	db.Init()
	defer db.Close()

	// Create the repository and service instances
	repo := db.NewDeliveryRepository(db.DB)
	fulfillmentService := service.NewFulfillmentService(repo)

	// Create the controller, passing in the service instance
	fulfillmentController := controller.NewFulfillmentController(fulfillmentService)

	// Set up gRPC server and register the controller as the service
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFulfillmentServiceServer(s, fulfillmentController) // Register the controller, not the service

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("shutting down gracefully...")
	s.GracefulStop()
}
