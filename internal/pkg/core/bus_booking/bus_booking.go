package bus_booking

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

const (
	BusBookingRouteMinLen = 4
	BusBookingRouteMaxLen = 10

	DateFormat = "2006-01-02"

	BusBookingMinSeatNumber = 1
	BusBookingMaxSeatNumber = 100
)

var (
	ErrValidate = errors.New("")

	ErrBusBookingNotExists     = errors.New("bus booking does not exist")
	ErrBusBookingAlreadyExists = errors.New("bus booking already exists")
	ErrRouteNameNotExist       = errors.New("route not exist")
	ErrInternal                = errors.New("internal error")
)

type Interface interface {
	List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error)
	Add(ctx context.Context, bb models.BusBooking) (uint, error)
	Get(ctx context.Context, id uint) (*models.BusBooking, error)
	ChangeSeat(ctx context.Context, id uint, newSeat uint) error
	ChangeDateSeat(ctx context.Context, id uint, date string, newSeat uint) error
	Delete(ctx context.Context, id uint) error
}

func New(repo repoPkg.Interface) Interface {
	return &core{
		repo: repo,
	}
}

type core struct {
	repo repoPkg.Interface
}

func (c *core) List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error) {
	result, err := c.repo.List(ctx, offset, limit)
	return result, repoErrorToBbError(err)
}

func (c *core) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	if err := checkCorrectRoute(bb.Route); err != nil {
		return 0, err
	}
	if err := checkCorrectDate(bb.Date); err != nil {
		return 0, err
	}
	if err := checkCorrectSeat(bb.Seat); err != nil {
		return 0, err
	}
	result, err := c.repo.Add(ctx, bb)
	return result, repoErrorToBbError(err)
}

func (c *core) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	result, err := c.repo.Get(ctx, id)
	return result, repoErrorToBbError(err)
}

func (c *core) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	if err := checkCorrectSeat(newSeat); err != nil {
		return err
	}
	return repoErrorToBbError(c.repo.ChangeSeat(ctx, id, newSeat))
}

func (c *core) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	if err := checkCorrectDate(newDate); err != nil {
		return err
	}
	if err := checkCorrectSeat(newSeat); err != nil {
		return err
	}
	return repoErrorToBbError(c.repo.ChangeDateSeat(ctx, id, newDate, newSeat))
}

func (c *core) Delete(ctx context.Context, id uint) error {
	return repoErrorToBbError(c.repo.Delete(ctx, id))
}

func checkCorrectRoute(route string) error {
	if len(route) < BusBookingRouteMinLen || len(route) > BusBookingRouteMaxLen {
		return errors.Wrapf(
			ErrValidate,
			"expected route length from %d to %d, got %d: [%v]",
			BusBookingRouteMinLen, BusBookingRouteMaxLen, len(route), route,
		)
	}
	return nil
}

func checkCorrectDate(dateString string) error {
	if _, err := time.Parse(DateFormat, dateString); err != nil {
		return errors.Wrapf(
			ErrValidate,
			"expected correct date in format [%v], got [%v]",
			DateFormat, dateString,
		)
	}
	return nil
}

func checkCorrectSeat(seat uint) error {
	if seat < BusBookingMinSeatNumber || seat > BusBookingMaxSeatNumber {
		return errors.Wrapf(
			ErrValidate,
			"expected seat number from %d to %d, got %d",
			BusBookingMinSeatNumber, BusBookingMaxSeatNumber, seat,
		)
	}
	return nil
}

func repoErrorToBbError(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, repoPkg.ErrRepoBusBookingNotExists):
		return ErrBusBookingNotExists
	case errors.Is(err, repoPkg.ErrRepoBusBookingAlreadyExists):
		return ErrBusBookingAlreadyExists
	case errors.Is(err, repoPkg.ErrRepoRouteNameNotExist):
		return ErrRouteNameNotExist
	case errors.Is(err, repoPkg.ErrRepoInternal):
		return ErrInternal
	}

	log.Printf("bus_booking::repoErrorToBbError unexpected error %s", err.Error())
	return ErrInternal
}
