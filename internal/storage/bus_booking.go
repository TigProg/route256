package storage

import (
	"fmt"
	"time"
)

const dateFormat = "2006-01-02"

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
	if len(route) < 4 || len(route) > 10 {
		return fmt.Errorf("bad route <%v>", route)
	}
	bb.route = route
	return nil
}

func (bb *BusBooking) SetDate(dateString string) error {
	if _, err := time.Parse(dateFormat, dateString); err != nil {
		return fmt.Errorf("bad date <%v>", dateString)
	}
	bb.date = dateString
	return nil
}

func (bb *BusBooking) SetSeat(seat uint) error {
	if seat == 0 || seat > 100 {
		return fmt.Errorf("bad seat <%d>", seat)
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
