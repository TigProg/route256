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
	t.Run("fail::validation_error", func(t *testing.T) {
		// Arrange
		// Act
		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: "spbufa",
			Date:  "2020-01-01 10:00:00",
			Seat:  uint32(25),
		}
		response, err := BusBookingClient.BusBookingAdd(ctx, &request)

		// Assert
		assert.Contains(t, err.Error(), "expected correct date in format [2006-01-02], got [2020-01-01 10:00:00]")
		assert.Nil(t, response)
	})
}
