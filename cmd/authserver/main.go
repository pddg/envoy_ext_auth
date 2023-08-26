package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"envoy-ext-auth/pkg/api"
)

func listenAndServe() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		return fmt.Errorf("net.Listen: %v", err)
	}
	grpcserver := api.NewServer()
	go func() {
		<-ctx.Done()
		shutdownCh := make(chan struct{})
		go func() {
			grpcserver.GracefulStop()
			close(shutdownCh)
		}()
		select {
		case <-time.After(5 * time.Second):
			fmt.Fprintf(os.Stderr, "Error: failed to stop grpc server gracefully")
			grpcserver.Stop()
		case <-shutdownCh:
		}
	}()
	return grpcserver.Serve(listener)
}

func main() {
	if err := listenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
