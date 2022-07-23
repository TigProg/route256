package storage

import (
	"log"
	"strconv"

	"github.com/pkg/errors"

	"gitlab.ozon.dev/tigprog/homeword-1/internal/tools"
)

var data map[uint]*BusBooking

var BusBookingNotExists = errors.New("bus booking dows not exist")
var BusBookingExists = errors.New("bus booking exists")

var EmptyField = errors.New("field is empty")
var FieldNotExist = errors.New("field does not exists")

func init() {
	log.Println("init storage")
	data = make(map[uint]*BusBooking)
	bb, _ := NewBusBooking("ABCD45", "2022-01-01", uint(7))
	if err := Add(bb); err != nil {
		log.Panic(err)
	}
}

func List() []*BusBooking {
	res := make([]*BusBooking, 0, len(data))
	for _, v := range data {
		res = append(res, v)
	}
	return res
}

func Get(id uint) (*BusBooking, error) {
	if bb, ok := data[id]; ok {
		return bb, nil
	}
	return nil, errors.Wrap(BusBookingNotExists, strconv.FormatUint(uint64(id), 10))
}

func Add(bb *BusBooking) error {
	if _, ok := data[bb.GetId()]; ok {
		return errors.Wrap(BusBookingExists, strconv.FormatUint(uint64(bb.GetId()), 10))
	}
	data[bb.GetId()] = bb
	return nil
}

func Update(id uint, field, newValue string) error {
	bb, ok := data[id]
	if !ok {
		return errors.Wrap(BusBookingNotExists, strconv.FormatUint(uint64(id), 10))
	}

	if field == "" {
		return EmptyField
	}
	switch field {
	case "route":
		return bb.SetRoute(newValue)
	case "date":
		return bb.SetDate(newValue)
	case "seat":
		seat, err := tools.StringToUint(newValue)
		if err != nil {
			return err
		}
		return bb.SetSeat(seat)
	default:
		return errors.Wrap(FieldNotExist, field)
	}
}

func Delete(id uint) error {
	if _, ok := data[id]; ok {
		delete(data, id)
		return nil
	}
	return errors.Wrap(BusBookingNotExists, strconv.FormatUint(uint64(id), 10))
}
