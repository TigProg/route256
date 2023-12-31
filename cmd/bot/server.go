package main

import (
	"context"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	apiPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/api"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runGRPCServer(ctx context.Context, bb bbPkg.Interface) { // TODO
	listener, err := net.Listen("tcp", configPkg.GRPCServerAddress)
	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(bb))

	if err = grpcServer.Serve(listener); err != nil {
		log.Panic(err)
	}
}

func runREST(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, configPkg.GRPCServerAddress, opts); err != nil {
		log.Panic(err)
	}

	if err := http.ListenAndServe(configPkg.RESTServerAddress, mux); err != nil {
		log.Panic(err)
	}
}
