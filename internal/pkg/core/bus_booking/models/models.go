package models

import "encoding/json"

type BusBooking struct {
	Id    uint
	Route string
	Date  string
	Seat  uint
}

func (bb BusBooking) MarshalBinary() ([]byte, error) {
	return json.Marshal(bb)
}

func (bb *BusBooking) UnmarshalBinary(bbByte []byte) error {
	return json.Unmarshal(bbByte, bb)
}
