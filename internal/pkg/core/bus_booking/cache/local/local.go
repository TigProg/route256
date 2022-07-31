package local

import (
	"sync"

	"github.com/pkg/errors"
	cachePkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/cache"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
)

var (
	ErrBusBookingNotExists     = errors.New("bus booking does not exist")
	ErrBusBookingAlreadyExists = errors.New("bus booking already exists")
)

func New() cachePkg.Interface {
	return &cache{
		mu:     sync.RWMutex{},
		nextId: 1,
		data:   map[uint]*models.BusBooking{},
	}
}

type cache struct {
	mu     sync.RWMutex
	nextId uint
	data   map[uint]*models.BusBooking
}

func (c *cache) List() []models.BusBooking {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]models.BusBooking, 0, len(c.data))
	for _, bb := range c.data {
		result = append(result, *bb)
	}
	return result
}

func (c *cache) Add(bb models.BusBooking) (uint, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if existedId, err := c.reverseSearch(bb.Route, bb.Date, bb.Seat); err == nil {
		return 0, errors.Wrapf(ErrBusBookingAlreadyExists, "%d", existedId)
	}

	var id = c.nextId
	bb.Id = id
	c.data[id] = &bb

	c.nextId++
	return id, nil
}

func (c *cache) Get(id uint) (*models.BusBooking, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if bb, ok := c.data[id]; ok {
		return bb, nil
	}
	return nil, errors.Wrapf(ErrBusBookingNotExists, "%d", id)
}

func (c *cache) ChangeSeat(id uint, newSeat uint) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	bb, ok := c.data[id]
	if !ok {
		return errors.Wrapf(ErrBusBookingNotExists, "%d", id)
	}
	if bb.Seat == newSeat {
		return nil // for idempotency
	}

	if existedId, err := c.reverseSearch(bb.Route, bb.Date, newSeat); err == nil {
		return errors.Wrapf(ErrBusBookingAlreadyExists, "%d", existedId)
	}
	bb.Seat = newSeat
	return nil
}

func (c *cache) ChangeDateSeat(id uint, newDate string, newSeat uint) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	bb, ok := c.data[id]
	if !ok {
		return errors.Wrapf(ErrBusBookingNotExists, "%d", id)
	}
	if bb.Seat == newSeat && bb.Date == newDate {
		return nil // for idempotency
	}

	if existedId, err := c.reverseSearch(bb.Route, newDate, newSeat); err == nil {
		return errors.Wrapf(ErrBusBookingAlreadyExists, "%d", existedId)
	}
	bb.Seat = newSeat
	bb.Date = newDate
	return nil
}

func (c *cache) Delete(id uint) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.data[id]; ok {
		delete(c.data, id)
		return nil
	}
	return errors.Wrapf(ErrBusBookingNotExists, "%d", id)
}

// reverseSearch - not thread-safe
func (c *cache) reverseSearch(route string, date string, seat uint) (uint, error) {
	for _, bb := range c.data {
		if bb.Route == route && bb.Date == date && bb.Seat == seat {
			return bb.Id, nil
		}
	}
	return 0, ErrBusBookingNotExists
}
