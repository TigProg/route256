package main

import (
	"context"
	"log"

	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conns, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewAdminClient(conns)

	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "custom", "hello")

	// test gRPC

	var responseAdd *pb.BusBookingAddResponse
	var responseList *pb.BusBookingListResponse

	responseAdd, err = client.BusBookingAdd(ctx, &pb.BusBookingAddRequest{
		Route: "aaaa", Date: "2022-01-01", Seat: 1,
	})
	if err != nil {
		panic(err)
	}
	log.Println("response: ", responseAdd)

	responseAdd, err = client.BusBookingAdd(ctx, &pb.BusBookingAddRequest{
		Route: "bbbb", Date: "2022-01-02", Seat: 2,
	})
	if err != nil {
		panic(err)
	}
	log.Println("response: ", responseAdd)

	responseAdd, err = client.BusBookingAdd(ctx, &pb.BusBookingAddRequest{
		Route: "cccc", Date: "2022-01-03", Seat: 3,
	})
	if err != nil {
		panic(err)
	}
	log.Println("response: ", responseAdd)

	responseAdd, err = client.BusBookingAdd(ctx, &pb.BusBookingAddRequest{
		Route: "dddd", Date: "2022-01-04", Seat: 4,
	})
	if err != nil {
		panic(err)
	}
	log.Println("response: ", responseAdd)

	responseList, err = client.BusBookingList(ctx, &pb.BusBookingListRequest{})
	if err != nil {
		panic(err)
	}
	log.Println("response: ", responseList)
}
