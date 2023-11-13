package grpc_repo

import (
	"context"

	log "github.com/sirupsen/logrus"

	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(client pb.AdminClient) repoPkg.Interface {
	return &repo{client}
}

type repo struct {
	client pb.AdminClient
}

func (r *repo) List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error) {
	request := pb.BusBookingListRequest{
		Offset: uint32(offset),
		Limit:  uint32(limit),
	}
	response, err := r.client.BusBookingList(ctx, &request)
	if err != nil {
		return nil, statusErrorToRepoError(err)
	}

	responseBbs := response.GetBusBookings()
	bbs := make([]models.BusBooking, 0, len(responseBbs))
	for _, responseBb := range responseBbs {
		bbs = append(bbs, models.BusBooking{
			Id:    uint(responseBb.Id),
			Route: responseBb.Route,
			Date:  responseBb.Date,
			Seat:  uint(responseBb.Seat),
		})
	}
	return bbs, nil
}

func (r *repo) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	request := pb.BusBookingAddRequest{
		Route: bb.Route,
		Date:  bb.Date,
		Seat:  uint32(bb.Seat),
	}
	response, err := r.client.BusBookingAdd(ctx, &request)
	if err != nil {
		return 0, statusErrorToRepoError(err)
	}
	return uint(response.GetId()), nil
}

func (r *repo) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	request := pb.BusBookingGetRequest{
		Id: uint32(id),
	}
	response, err := r.client.BusBookingGet(ctx, &request)
	if err != nil {
		return nil, statusErrorToRepoError(err)
	}

	responseBb := response.GetBusBooking()
	return &models.BusBooking{
		Id:    uint(responseBb.Id),
		Route: responseBb.Route,
		Date:  responseBb.Date,
		Seat:  uint(responseBb.Seat),
	}, nil
}

func (r *repo) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	request := pb.BusBookingChangeSeatRequest{
		Id:   uint32(id),
		Seat: uint32(newSeat),
	}
	_, err := r.client.BusBookingChangeSeat(ctx, &request)
	if err != nil {
		return statusErrorToRepoError(err)
	}
	return nil
}

func (r *repo) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	request := pb.BusBookingChangeDateSeatRequest{
		Id:   uint32(id),
		Date: newDate,
		Seat: uint32(newSeat),
	}
	_, err := r.client.BusBookingChangeDateSeat(ctx, &request)
	if err != nil {
		return statusErrorToRepoError(err)
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, id uint) error {
	request := pb.BusBookingDeleteRequest{
		Id: uint32(id),
	}
	_, err := r.client.BusBookingDelete(ctx, &request)
	if err != nil {
		return statusErrorToRepoError(err)
	}
	return nil
}

func statusErrorToRepoError(err error) error {
	st, ok := status.FromError(err)
	if !ok { // network error
		log.Errorf("grpc_repo::statusErrorToRepoError not status error %s", err.Error())
		return repoPkg.ErrRepoInternal
	}

	switch st.Code() {
	case codes.NotFound:
		switch st.Message() {
		case repoPkg.ErrRepoBusBookingNotExists.Error():
			return repoPkg.ErrRepoBusBookingNotExists
		case repoPkg.ErrRepoRouteNameNotExist.Error():
			return repoPkg.ErrRepoRouteNameNotExist
		default:
			log.Errorf("grpc_repo::statusErrorToRepoError unexpected not found error %s", err.Error())
			return repoPkg.ErrRepoInternal
		}
	case codes.AlreadyExists:
		return repoPkg.ErrRepoBusBookingAlreadyExists
	case codes.Internal:
		return repoPkg.ErrRepoInternal
	}

	log.Errorf("grpc_repo::statusErrorToRepoError unexpected error %s", err.Error())
	return repoPkg.ErrRepoInternal
}
