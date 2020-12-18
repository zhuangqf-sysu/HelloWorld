package server

import (
	"log"
	"net"

	"HelloWorld/api"
	"HelloWorld/internal/service"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewSever(service *service.Service) (*Server, func(), error) {
	grpcServer := grpc.NewServer()
	api.RegisterHelloWorldServer(grpcServer, &GrpcServerImpl{service: service})

	server := &Server{grpcServer: grpcServer}

	return server, server.Close, nil
}

func (server *Server) Start() error {
	// todo from config
	addr := "0.0.0.0:8000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "start grpc server err")
	}
	log.Printf("start grpc server on %s\n", addr)
	return server.grpcServer.Serve(listener)
}

func (server *Server) Stop() {
	server.grpcServer.Stop()
}

func (server *Server) Close() {
	if server.grpcServer != nil {
		server.grpcServer.Stop()
	}
}
