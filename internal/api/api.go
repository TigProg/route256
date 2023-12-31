package api

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
)

func New(busBooking bbPkg.Interface) *Implementation {
	return &Implementation{
		busBooking: busBooking,
	}
}

type Implementation struct {
	pb.UnimplementedAdminServer
	busBooking bbPkg.Interface
}

func (i *Implementation) BusBookingList(ctx context.Context, in *pb.BusBookingListRequest) (*pb.BusBookingListResponse, error) {
	offset := uint(in.GetOffset())
	limit := uint(in.GetLimit())

	bbs, err := i.busBooking.List(ctx, offset, limit)
	if err != nil {
		return nil, bbErrorToStatusError(err)
	}

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

func (i *Implementation) BusBookingAdd(ctx context.Context, in *pb.BusBookingAddRequest) (*pb.BusBookingAddResponse, error) {
	id, err := i.busBooking.Add(ctx, models.BusBooking{
		Id:    0,
		Route: in.GetRoute(),
		Date:  in.GetDate(),
		Seat:  uint(in.GetSeat()),
	})
	if err != nil {
		return nil, bbErrorToStatusError(err)
	}
	return &pb.BusBookingAddResponse{Id: uint32(id)}, nil
}

func (i *Implementation) BusBookingGet(ctx context.Context, in *pb.BusBookingGetRequest) (*pb.BusBookingGetResponse, error) {
	bb, err := i.busBooking.Get(ctx, uint(in.GetId()))
	if err != nil {
		return nil, bbErrorToStatusError(err)
	}
	return &pb.BusBookingGetResponse{BusBooking: &pb.BusBooking{
		Id:    uint32(bb.Id),
		Route: bb.Route,
		Date:  bb.Date,
		Seat:  uint32(bb.Seat),
	}}, nil
}

func (i *Implementation) BusBookingChangeSeat(ctx context.Context, in *pb.BusBookingChangeSeatRequest) (*pb.BusBookingChangeSeatResponse, error) {
	err := i.busBooking.ChangeSeat(
		ctx,
		uint(in.GetId()),
		uint(in.GetSeat()),
	)
	if err != nil {
		return nil, bbErrorToStatusError(err)
	}
	return &pb.BusBookingChangeSeatResponse{}, nil
}

func (i *Implementation) BusBookingChangeDateSeat(ctx context.Context, in *pb.BusBookingChangeDateSeatRequest) (*pb.BusBookingChangeDateSeatResponse, error) {
	err := i.busBooking.ChangeDateSeat(
		ctx,
		uint(in.GetId()),
		in.GetDate(),
		uint(in.GetSeat()),
	)
	if err != nil {
		return nil, bbErrorToStatusError(err)
	}
	return &pb.BusBookingChangeDateSeatResponse{}, nil
}

func (i *Implementation) BusBookingDelete(ctx context.Context, in *pb.BusBookingDeleteRequest) (*pb.BusBookingDeleteResponse, error) {
	err := i.busBooking.Delete(ctx, uint(in.GetId()))
	if err != nil {
		return nil, bbErrorToStatusError(err)
	}
	return &pb.BusBookingDeleteResponse{}, nil
}

func bbErrorToStatusError(err error) error {
	switch {
	case errors.Is(err, bbPkg.ErrValidate):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, bbPkg.ErrBusBookingNotExists):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, bbPkg.ErrBusBookingAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, bbPkg.ErrRouteNameNotExist):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, bbPkg.ErrInternal):
		return status.Error(codes.Internal, err.Error())
	}

	log.Errorf("api::bbErrorToStatusError unexpected error %s", err.Error())
	return status.Error(codes.Internal, err.Error())
}
