package api

import (
	"context"

	authv3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type AuthHandler struct {
	authv3.UnimplementedAuthorizationServer
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Check(ctx context.Context, in *authv3.CheckRequest) (*authv3.CheckResponse, error) {
	return &authv3.CheckResponse{
		Status: &status.Status{
			Code:    int32(code.Code_PERMISSION_DENIED),
			Message: "all requests are denied",
		},
	}, nil
}
