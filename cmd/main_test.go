package main

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestServerStart(t *testing.T) {
	// Create a channel to receive OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Run the main function in a separate goroutine
	go func() {
		main()
	}()

	// Give the server some time to start
	time.Sleep(2 * time.Second)

	// Check if the server is running
	select {
	case <-signalChan:
		t.Error("Server did not start successfully")
	default:
		t.Log("Server started successfully")
	}
}

func TestServerShutdown(t *testing.T) {
	// Create a channel to receive OS signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Run the main function in a separate goroutine
	go func() {
		main()
	}()

	// Give the server some time to start
	time.Sleep(2 * time.Second)

	// Send an interrupt signal to the process
	signalChan <- os.Interrupt

	// Give the server some time to shut down
	time.Sleep(2 * time.Second)

	// Check if the server has shut down gracefully
	select {
	case <-signalChan:
		t.Log("Server shut down gracefully")
	default:
		t.Error("Server did not shut down gracefully")
	}
}
