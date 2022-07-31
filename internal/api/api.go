package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
)

func New(busBooking bbPkg.Interface) pb.AdminServer {
	return &implementation{
		busBooking: busBooking,
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	busBooking bbPkg.Interface
}

func (i *implementation) BusBookingList(ctx context.Context, in *pb.BusBookingListRequest) (*pb.BusBookingListResponse, error) {
	bbs := i.busBooking.List()

	result := make([]*pb.BusBooking, 0, len(bbs))
	for _, bb := range bbs {
		result = append(result, &pb.BusBooking{
			Id:    uint32(bb.Id),
			Route: bb.Route,
			Date:  bb.Date,
			Seat:  uint32(bb.Seat),
		})
	}
	return &pb.BusBookingListResponse{
		BusBookings: result,
	}, nil
}

func (i *implementation) BusBookingAdd(ctx context.Context, in *pb.BusBookingAddRequest) (*pb.BusBookingAddResponse, error) {
	id, err := i.busBooking.Add(models.BusBooking{
		Id:    0,
		Route: in.GetRoute(),
		Date:  in.GetDate(),
		Seat:  uint(in.GetSeat()),
	})
	if err != nil { // TODO enrich work with errors
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.BusBookingAddResponse{Id: uint32(id)}, nil
}

func (i *implementation) BusBookingGet(ctx context.Context, in *pb.BusBookingGetRequest) (*pb.BusBookingGetResponse, error) {
	bb, err := i.busBooking.Get(uint(in.GetId()))
	if err != nil { // TODO enrich work with errors
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.BusBookingGetResponse{BusBooking: &pb.BusBooking{
		Id:    uint32(bb.Id),
		Route: bb.Route,
		Date:  bb.Date,
		Seat:  uint32(bb.Seat),
	}}, nil
}

func (i *implementation) BusBookingChangeSeat(ctx context.Context, in *pb.BusBookingChangeSeatRequest) (*pb.BusBookingChangeSeatResponse, error) {
	err := i.busBooking.ChangeSeat(
		uint(in.GetId()),
		uint(in.GetSeat()),
	)
	if err != nil { // TODO enrich work with errors
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.BusBookingChangeSeatResponse{}, nil
}

func (i *implementation) BusBookingChangeDateSeat(ctx context.Context, in *pb.BusBookingChangeDateSeatRequest) (*pb.BusBookingChangeDateSeatResponse, error) {
	err := i.busBooking.ChangeDateSeat(
		uint(in.GetId()),
		in.GetDate(),
		uint(in.GetSeat()),
	)
	if err != nil { // TODO enrich work with errors
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.BusBookingChangeDateSeatResponse{}, nil
}

func (i *implementation) BusBookingDelete(ctx context.Context, in *pb.BusBookingDeleteRequest) (*pb.BusBookingDeleteResponse, error) {
	err := i.busBooking.Delete(uint(in.GetId()))
	if err != nil { // TODO enrich work with errors
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.BusBookingDeleteResponse{}, nil
}
