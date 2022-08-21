//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
)

func TestAddBusBookingIntegration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		const (
			bbRoute = "ufaspb"
			bbDate  = "2025-01-01"
			bbSeat  = uint(2)
		)

		// Act
		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: bbRoute,
			Date:  bbDate,
			Seat:  uint32(bbSeat),
		}
		_, err := BusBookingClient.BusBookingAdd(ctx, &request)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("fail::invalid_route", func(t *testing.T) {
		// Arrange
		const (
			bbRoute = "abcdefg"
			bbDate  = "2025-01-01"
			bbSeat  = uint(2)
		)

		// Act
		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: bbRoute,
			Date:  bbDate,
			Seat:  uint32(bbSeat),
		}
		_, err := BusBookingClient.BusBookingAdd(ctx, &request)

		// Assert
		assert.Contains(t, err.Error(), "route not exist")
	})
}
