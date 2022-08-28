package kafka

import (
	"errors"

	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
)

var ErrIncorrectKeyMessage = errors.New("incorrect key in message") // TODO

type CommonMessage struct {
	// trace
	Key   string
	Id    uint
	Route string
	Date  string
	Seat  uint
}

func (msg *CommonMessage) ToSpecificMessage() (SpecificMessage, error) {
	switch msg.Key {
	case "add":
		return AddMessage{Bb: models.BusBooking{
			Id:    msg.Id,
			Route: msg.Route,
			Date:  msg.Date,
			Seat:  msg.Seat,
		}}, nil
	case "change_seat":
		return ChangeSeatMessage{
			Id:      msg.Id,
			NewSeat: msg.Seat,
		}, nil
	case "change_date_seat":
		return &ChangeDateSeatMessage{
			Id:      msg.Id,
			NewDate: msg.Date,
			NewSeat: msg.Seat,
		}, nil
	case "delete":
		return DeleteMessage{
			Id: msg.Id,
		}, nil
	default:
		return nil, ErrIncorrectKeyMessage
	}
}

type SpecificMessage interface {
	ToCommon() CommonMessage
}

type AddMessage struct {
	Bb models.BusBooking
}

func (msg AddMessage) ToCommon() CommonMessage {
	return CommonMessage{
		Key:   "add",
		Id:    msg.Bb.Id,
		Route: msg.Bb.Route,
		Date:  msg.Bb.Date,
		Seat:  msg.Bb.Seat,
	}
}

type ChangeSeatMessage struct {
	Id      uint
	NewSeat uint
}

func (msg ChangeSeatMessage) ToCommon() CommonMessage {
	return CommonMessage{
		Id:   msg.Id,
		Seat: msg.NewSeat,
	}
}

type ChangeDateSeatMessage struct {
	Id      uint
	NewDate string
	NewSeat uint
}

func (msg ChangeDateSeatMessage) ToCommon() CommonMessage {
	return CommonMessage{
		Id:   msg.Id,
		Date: msg.NewDate,
		Seat: msg.NewSeat,
	}
}

type DeleteMessage struct {
	Id uint
}

func (msg DeleteMessage) ToCommon() CommonMessage {
	return CommonMessage{
		Id: msg.Id,
	}
}
