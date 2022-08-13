package local

import (
	"context"
	"sort"
	"sync"

	"github.com/pkg/errors"
	configPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/config"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

var routeAllowList map[string]struct{}

func init() {
	// only for local test
	// copy of second migration
	routeAllowList = make(map[string]struct{})
	routeAllowList["ufaspb"] = struct{}{}
	routeAllowList["spbufa"] = struct{}{}
	routeAllowList["ufamsk"] = struct{}{}
	routeAllowList["mskufa"] = struct{}{}
	routeAllowList["spbmsk"] = struct{}{}
	routeAllowList["mskspb"] = struct{}{}
	routeAllowList["ufanowhere"] = struct{}{}
	routeAllowList["nowhereufa"] = struct{}{}
}

func New() repoPkg.Interface {
	return &local{
		mu:     sync.RWMutex{},
		nextId: 1,
		data:   map[uint]*models.BusBooking{},
		poolCh: make(chan struct{}, configPkg.LocalCachePoolSize),
	}
}

type local struct {
	mu     sync.RWMutex
	nextId uint
	data   map[uint]*models.BusBooking
	poolCh chan struct{}
}

func (c *local) List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
	}()

	result := make([]models.BusBooking, 0, len(c.data))
	for _, bb := range c.data {
		result = append(result, *bb)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Id < result[j].Id
	})

	leftBound := min(int(offset), len(result))
	rightBound := min(int(offset+limit), len(result))

	return result[leftBound:rightBound], nil
}

func (c *local) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	if _, found := routeAllowList[bb.Route]; !found {
		return 0, repoPkg.ErrRouteNameNotExist
	}

	if existedId, err := c.reverseSearch(bb.Route, bb.Date, bb.Seat); err == nil {
		return 0, errors.Wrapf(repoPkg.ErrBusBookingAlreadyExists, "%d", existedId)
	}

	var id = c.nextId
	bb.Id = id
	c.data[id] = &bb

	c.nextId++
	return id, nil
}

func (c *local) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	c.poolCh <- struct{}{}
	c.mu.RLock()
	defer func() {
		c.mu.RUnlock()
		<-c.poolCh
	}()

	if bb, ok := c.data[id]; ok {
		return bb, nil
	}
	return nil, errors.Wrapf(repoPkg.ErrBusBookingNotExists, "%d", id)
}

func (c *local) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	bb, ok := c.data[id]
	if !ok {
		return errors.Wrapf(repoPkg.ErrBusBookingNotExists, "%d", id)
	}
	if bb.Seat == newSeat {
		return nil // for idempotency
	}

	if existedId, err := c.reverseSearch(bb.Route, bb.Date, newSeat); err == nil {
		return errors.Wrapf(repoPkg.ErrBusBookingAlreadyExists, "%d", existedId)
	}
	bb.Seat = newSeat
	return nil
}

func (c *local) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	bb, ok := c.data[id]
	if !ok {
		return errors.Wrapf(repoPkg.ErrBusBookingNotExists, "%d", id)
	}
	if bb.Seat == newSeat && bb.Date == newDate {
		return nil // for idempotency
	}

	if existedId, err := c.reverseSearch(bb.Route, newDate, newSeat); err == nil {
		return errors.Wrapf(repoPkg.ErrBusBookingAlreadyExists, "%d", existedId)
	}
	bb.Seat = newSeat
	bb.Date = newDate
	return nil
}

func (c *local) Delete(ctx context.Context, id uint) error {
	c.poolCh <- struct{}{}
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		<-c.poolCh
	}()

	if _, ok := c.data[id]; ok {
		delete(c.data, id)
		return nil
	}
	return errors.Wrapf(repoPkg.ErrBusBookingNotExists, "%d", id)
}

// reverseSearch - not thread-safe
func (c *local) reverseSearch(route string, date string, seat uint) (uint, error) {
	for _, bb := range c.data {
		if bb.Route == route && bb.Date == date && bb.Seat == seat {
			return bb.Id, nil
		}
	}
	return 0, repoPkg.ErrBusBookingNotExists
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
