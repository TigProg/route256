package local

// TODO add mutex

import (
	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/cache"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
)

var (
	ErrBusBookingNotExists = errors.New("bus booking does not exist")
	//ErrBusBookingExists    = errors.New("bus booking exists")
)

func New() cachePkg.Interface {
	return &cache{
		nextId: 1,
		data:   map[uint]*models.BusBooking{},
	}
}

type cache struct {
	nextId uint
	data   map[uint]*models.BusBooking
}

func (c *cache) List() []models.BusBooking {
	result := make([]models.BusBooking, 0, len(c.data))
	for _, bb := range c.data {
		result = append(result, *bb)
	}
	return result
}

func (c *cache) Add(bb models.BusBooking) (uint, error) {
	// TODO add reverse search

	var id = c.nextId
	bb.Id = id
	c.data[id] = &bb
	c.nextId++
	return id, nil
}

func (c *cache) Get(id uint) (*models.BusBooking, error) {
	if bb, ok := c.data[id]; ok {
		return bb, nil
	}
	return nil, errors.Wrapf(ErrBusBookingNotExists, "%d", id)
}

func (c *cache) ChangeSeat(id uint, newSeat uint) error {
	bb, ok := c.data[id]
	if !ok {
		return errors.Wrapf(ErrBusBookingNotExists, "%d", id)
	}
	if bb.Seat == newSeat {
		return nil // for idempotency
	}

	// TODO add reverse search
	bb.Seat = newSeat
	return nil
}

func (c *cache) ChangeDateSeat(id uint, newDate string, newSeat uint) error {
	bb, ok := c.data[id]
	if !ok {
		return errors.Wrapf(ErrBusBookingNotExists, "%d", id)
	}
	if bb.Seat == newSeat && bb.Date == newDate {
		return nil // for idempotency
	}

	// TODO add reverse search
	bb.Seat = newSeat
	bb.Date = newDate
	return nil
}

func (c *cache) Delete(id uint) error {
	if _, ok := c.data[id]; ok {
		delete(c.data, id)
		return nil
	}
	return errors.Wrapf(ErrBusBookingNotExists, "%d", id)
}
