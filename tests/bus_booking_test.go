//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddBusBookingIntegration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: "ufaspb",
			Date:  "2025-01-01",
			Seat:  1,
		}

		// Act
		_, err := BusBookingClient.BusBookingAdd(ctx, &request)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("fail::incorrect_route", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: "a",
			Date:  "2025-01-01",
			Seat:  1,
		}

		// Act
		_, err := BusBookingClient.BusBookingAdd(ctx, &request)
		st, ok := status.FromError(err)

		// Assert
		require.True(t, ok)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Contains(t, st.Message(), "expected route length")

	})
}
