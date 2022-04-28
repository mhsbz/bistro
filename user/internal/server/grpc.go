package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/mhsbz/bistro/user/internal/services"
	v1 "github.com/mhsbz/uinterface/user/v1"
)

func NewGrpcServer(userService *services.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterUserServer(srv, userService)
	return srv
}
