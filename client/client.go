package main

import (
	"context"
	"log"

	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conns, err := grpc.Dial(configPkg.GRPCClientTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}

	client := pb.NewAdminClient(conns)

	ctx := context.Background()

	// test gRPC

	var responseAdd *pb.BusBookingAddResponse
	var responseList *pb.BusBookingListResponse

	responseAdd, err = client.BusBookingAdd(ctx, &pb.BusBookingAddRequest{
		Route: "aaaa", Date: "2022-01-01", Seat: 1,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Println("response: ", responseAdd)

	responseAdd, err = client.BusBookingAdd(ctx, &pb.BusBookingAddRequest{
		Route: "bbbb", Date: "2022-01-02", Seat: 2,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Println("response: ", responseAdd)

	responseAdd, err = client.BusBookingAdd(ctx, &pb.BusBookingAddRequest{
		Route: "cccc", Date: "2022-01-03", Seat: 3,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Println("response: ", responseAdd)

	responseAdd, err = client.BusBookingAdd(ctx, &pb.BusBookingAddRequest{
		Route: "dddd", Date: "2022-01-04", Seat: 4,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Println("response: ", responseAdd)

	responseList, err = client.BusBookingList(ctx, &pb.BusBookingListRequest{})
	if err != nil {
		log.Panic(err)
	}
	log.Println("response: ", responseList)
}
