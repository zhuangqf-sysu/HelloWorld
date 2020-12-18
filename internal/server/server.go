package server

import (
	"HelloWorld/api"
	"HelloWorld/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewSever(service *service.Service) (*Server, error) {
	grpcServer := grpc.NewServer()
	api.RegisterHelloWorldServer(grpcServer, &GrpcServerImpl{service: service})

	return &Server{
		grpcServer: grpcServer,
	}, nil
}

func (server *Server) Close() {
	if server.grpcServer != nil {
		server.grpcServer.Stop()
	}
}
