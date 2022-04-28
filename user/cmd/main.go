package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/mhsbz/bistro/user/internal/server"
	"github.com/mhsbz/bistro/user/internal/services"
	"os"
)

var (
	id, _   = os.Hostname()
	version = "1.0"
	name    = "user"
)

func newApp(srv *grpc.Server, logger log.Logger) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name("user"),
		kratos.Logger(logger),
		kratos.Server(srv),
	)
}

func main() {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", name,
		"service.version", version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	service := services.NewUserService()
	srv := server.NewGrpcServer(service)
	app := newApp(srv, logger)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
