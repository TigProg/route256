package bus_booking

import (
	"context"
	"fmt"
	"time"

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

type Interface interface {
	List(ctx context.Context) ([]models.BusBooking, error)
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

func (c *core) List(ctx context.Context) ([]models.BusBooking, error) {
	return c.repo.List(ctx)
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
	return c.repo.Add(ctx, bb)
}

func (c *core) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	return c.repo.Get(ctx, id)
}

func (c *core) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	if err := checkCorrectSeat(newSeat); err != nil {
		return err
	}
	return c.repo.ChangeSeat(ctx, id, newSeat)
}

func (c *core) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	if err := checkCorrectDate(newDate); err != nil {
		return err
	}
	if err := checkCorrectSeat(newSeat); err != nil {
		return err
	}
	return c.repo.ChangeDateSeat(ctx, id, newDate, newSeat)
}

func (c *core) Delete(ctx context.Context, id uint) error {
	return c.repo.Delete(ctx, id)
}

func checkCorrectRoute(route string) error {
	if len(route) < BusBookingRouteMinLen || len(route) > BusBookingRouteMaxLen {
		return fmt.Errorf(
			"expected route length from %d to %d, got %d: [%v]",
			BusBookingRouteMinLen, BusBookingRouteMaxLen, len(route), route,
		)
	}
	return nil
}

func checkCorrectDate(dateString string) error {
	if _, err := time.Parse(DateFormat, dateString); err != nil {
		return fmt.Errorf(
			"expected correct date in format [%v], got [%v]",
			DateFormat, dateString,
		)
	}
	return nil
}

func checkCorrectSeat(seat uint) error {
	if seat < BusBookingMinSeatNumber || seat > BusBookingMaxSeatNumber {
		return fmt.Errorf(
			"expected seat number from %d to %d, got %d",
			BusBookingMinSeatNumber, BusBookingMaxSeatNumber, seat,
		)
	}
	return nil
}
