package api

import (
	authv3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewServer() *grpc.Server {
	srv := grpc.NewServer()
	authv3.RegisterAuthorizationServer(srv, NewAuthHandler())
	reflection.Register(srv)
	return srv
}
