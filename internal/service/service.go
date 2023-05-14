package service

import (
	"auth/internal/config"
	"auth/internal/server"
	"fmt"
	"net"

	authpb "github.com/the-spine/spine-protos-go/auth"

	"google.golang.org/grpc"
)

func StartGrpcServer(config *config.Config) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Api.Host, config.Api.Port))
	if err != nil {
		return nil, err
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	authpb.RegisterAuthServiceServer(grpcServer, server.GetAuthServer())

	grpcServer.Serve(lis)

	return grpcServer, nil
}
