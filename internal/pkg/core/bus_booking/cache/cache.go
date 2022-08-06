package cache

import (
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
)

type Interface interface {
	List() []models.BusBooking
	Add(bb models.BusBooking) (uint, error)
	Get(id uint) (*models.BusBooking, error)
	ChangeSeat(id uint, newSeat uint) error
	ChangeDateSeat(id uint, newDate string, newSeat uint) error
	Delete(id uint) error
}
