package repository

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
)

var (
	ErrBusBookingNotExists     = errors.New("bus booking does not exist")
	ErrBusBookingAlreadyExists = errors.New("bus booking already exists")
	ErrRouteNameNotExist       = errors.New("route not exist")
)

type Interface interface {
	List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error)
	Add(ctx context.Context, bb models.BusBooking) (uint, error)
	Get(ctx context.Context, id uint) (*models.BusBooking, error)
	ChangeSeat(ctx context.Context, id uint, newSeat uint) error
	ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error
	Delete(ctx context.Context, id uint) error
}
