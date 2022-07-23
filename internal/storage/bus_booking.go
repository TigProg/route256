package storage

import (
	"fmt"
	"time"

	"gitlab.ozon.dev/tigprog/homeword-1/config"
)

var lastId = uint(0)

type BusBooking struct {
	id    uint
	route string
	date  string
	seat  uint
}

func NewBusBooking(route, dateString string, seat uint) (*BusBooking, error) {
	bb := BusBooking{}

	if err := bb.SetRoute(route); err != nil {
		return nil, err
	}
	if err := bb.SetDate(dateString); err != nil {
		return nil, err
	}
	if err := bb.SetSeat(seat); err != nil {
		return nil, err
	}
	lastId++
	bb.id = lastId
	return &bb, nil
}

func (bb *BusBooking) SetRoute(route string) error {
	if len(route) < config.BusBookingRouteMinLen || len(route) > config.BusBookingRouteMaxLen {
		return fmt.Errorf("expected route len from <%d> to <%d>, got <%d>: <%v>",
			config.BusBookingRouteMinLen, config.BusBookingRouteMaxLen, len(route), route)
	}
	bb.route = route
	return nil
}

func (bb *BusBooking) SetDate(dateString string) error {
	if _, err := time.Parse(config.DateFormat, dateString); err != nil {
		return fmt.Errorf("expected correct date in format <%v>, got <%v>", config.DateFormat, dateString)
	}
	bb.date = dateString
	return nil
}

func (bb *BusBooking) SetSeat(seat uint) error {
	if seat < config.BusBookingMinSeatNumber || seat > config.BusBookingMaxSeatNumber {
		return fmt.Errorf("expected seat number from <%d> to <%d>, got <%d>",
			config.BusBookingMinSeatNumber, config.BusBookingMaxSeatNumber, seat)
	}
	bb.seat = seat
	return nil
}

func (bb BusBooking) String() string {
	return fmt.Sprintf("%d: %s / %s / %d", bb.id, bb.route, bb.date, bb.seat)
}

func (bb BusBooking) GetRoute() string {
	return bb.route
}

func (bb BusBooking) GetDate() string {
	return bb.date
}

func (bb BusBooking) GetSeat() uint {
	return bb.seat
}

func (bb BusBooking) GetId() uint {
	return bb.id
}
