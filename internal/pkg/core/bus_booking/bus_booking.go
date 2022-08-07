package bus_booking

import (
	"context"
	"fmt"
	"time"

	cachePkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/cache"
	localCachePkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/cache/local"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
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

func New() Interface {
	return &core{
		cache: localCachePkg.New(),
	}
}

type core struct {
	cache cachePkg.Interface
}

func (c *core) List(ctx context.Context) ([]models.BusBooking, error) {
	return c.cache.List()
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
	return c.cache.Add(bb)
}

func (c *core) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	return c.cache.Get(id)
}

func (c *core) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	if err := checkCorrectSeat(newSeat); err != nil {
		return err
	}
	return c.cache.ChangeSeat(id, newSeat)
}

func (c *core) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	if err := checkCorrectDate(newDate); err != nil {
		return err
	}
	if err := checkCorrectSeat(newSeat); err != nil {
		return err
	}
	return c.cache.ChangeDateSeat(id, newDate, newSeat)
}

func (c *core) Delete(ctx context.Context, id uint) error {
	return c.cache.Delete(id)
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
