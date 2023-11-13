package api

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
)

func New(repo repoPkg.Interface) pb.AdminServer {
	return &implementation{
		repo: repo,
	}
}

type implementation struct {
	pb.UnimplementedAdminServer
	repo repoPkg.Interface
}

func (i *implementation) BusBookingList(ctx context.Context, in *pb.BusBookingListRequest) (*pb.BusBookingListResponse, error) {
	offset := uint(in.GetOffset())
	limit := uint(in.GetLimit())

	bbs, err := i.repo.List(ctx, offset, limit)
	if err != nil {
		return nil, repoErrorToStatusError(err)
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

func (i *implementation) BusBookingAdd(ctx context.Context, in *pb.BusBookingAddRequest) (*pb.BusBookingAddResponse, error) {
	id, err := i.repo.Add(ctx, models.BusBooking{
		Id:    0,
		Route: in.GetRoute(),
		Date:  in.GetDate(),
		Seat:  uint(in.GetSeat()),
	})
	if err != nil {
		return nil, repoErrorToStatusError(err)
	}
	return &pb.BusBookingAddResponse{Id: uint32(id)}, nil
}

func (i *implementation) BusBookingGet(ctx context.Context, in *pb.BusBookingGetRequest) (*pb.BusBookingGetResponse, error) {
	bb, err := i.repo.Get(ctx, uint(in.GetId()))
	if err != nil {
		return nil, repoErrorToStatusError(err)
	}
	return &pb.BusBookingGetResponse{BusBooking: &pb.BusBooking{
		Id:    uint32(bb.Id),
		Route: bb.Route,
		Date:  bb.Date,
		Seat:  uint32(bb.Seat),
	}}, nil
}

func (i *implementation) BusBookingChangeSeat(ctx context.Context, in *pb.BusBookingChangeSeatRequest) (*pb.BusBookingChangeSeatResponse, error) {
	err := i.repo.ChangeSeat(
		ctx,
		uint(in.GetId()),
		uint(in.GetSeat()),
	)
	if err != nil {
		return nil, repoErrorToStatusError(err)
	}
	return &pb.BusBookingChangeSeatResponse{}, nil
}

func (i *implementation) BusBookingChangeDateSeat(ctx context.Context, in *pb.BusBookingChangeDateSeatRequest) (*pb.BusBookingChangeDateSeatResponse, error) {
	err := i.repo.ChangeDateSeat(
		ctx,
		uint(in.GetId()),
		in.GetDate(),
		uint(in.GetSeat()),
	)
	if err != nil {
		return nil, repoErrorToStatusError(err)
	}
	return &pb.BusBookingChangeDateSeatResponse{}, nil
}

func (i *implementation) BusBookingDelete(ctx context.Context, in *pb.BusBookingDeleteRequest) (*pb.BusBookingDeleteResponse, error) {
	err := i.repo.Delete(ctx, uint(in.GetId()))
	if err != nil {
		return nil, repoErrorToStatusError(err)
	}
	return &pb.BusBookingDeleteResponse{}, nil
}

func repoErrorToStatusError(err error) error {
	switch {
	case errors.Is(err, repoPkg.ErrRepoBusBookingNotExists):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, repoPkg.ErrRepoBusBookingAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, repoPkg.ErrRepoRouteNameNotExist):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, repoPkg.ErrRepoInternal):
		return status.Error(codes.Internal, err.Error())
	}

	log.Errorf("data_api::repoErrorToStatusError unexpected error %s", err.Error())
	return status.Error(codes.Internal, err.Error())
}
