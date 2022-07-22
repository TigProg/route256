package storage

import (
	"log"
	"strconv"

	"github.com/pkg/errors"
)

var data map[uint]*BusBooking

//var BusBookingNotExists = errors.New("bus booking dows not exist")
var BusBookingExists = errors.New("bus booking exists")

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

func Add(bb *BusBooking) error {
	if _, ok := data[bb.GetId()]; ok {
		return errors.Wrap(BusBookingExists, strconv.FormatUint(uint64(bb.GetId()), 10))
	}
	data[bb.GetId()] = bb
	return nil
}

//func Update(u *User) error {
//	if _, ok := data[u.GetId()]; !ok {
//		return errors.Wrap(UserNotExists, strconv.FormatUint(uint64(u.GetId()), 10))
//	}
//	data[u.GetId()] = u
//	return nil
//}
//
//func Delete(id uint) error {
//	if _, ok := data[id]; ok {
//		delete(data, id)
//		return nil
//	}
//	return errors.Wrap(UserNotExists, strconv.FormatUint(uint64(id), 10))
//}
