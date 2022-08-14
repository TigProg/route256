package main

import (
	"context"
	"log"
	"net"

	dataApiPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/data_api"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func prepareRepoGRPCClient(address string) pb.AdminClient {
	conns, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	return pb.NewAdminClient(conns)
}

func runRepoGRPCServer(ctx context.Context, repo repoPkg.Interface, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, dataApiPkg.New(repo))

	if err = grpcServer.Serve(listener); err != nil {
		log.Panic(err)
	}
}
