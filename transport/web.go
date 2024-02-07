package web

import (
	"net/http"

	authService "github.com/antoha2/auth/service"
	pb "github.com/antoha2/auth/transport/protoAPI"
)

type Transport interface {
}

type webImpl struct {
	authService authService.AuthService
	server      *http.Server
}

func NewWeb(authService authService.AuthService) *webImpl {
	return &webImpl{
		authService: authService,
	}
}

type GRPCImpl struct {
	pb.UnimplementedTaskServiceServer
	authService *authService.AuthService
	//authService.AuthService
	//pb.TaskServiceServer
}

func NewGRPC(authService *authService.AuthService) *GRPCImpl {
	return &GRPCImpl{
		authService: authService,
	}
}
