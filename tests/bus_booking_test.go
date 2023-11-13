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

func TestListBusBookingIntegration(t *testing.T) {
	t.Run("success::empty", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingListRequest{
			Offset: 0,
			Limit:  100,
		}

		// Act
		response, err := BusBookingClient.BusBookingList(ctx, &request)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, 0, len(response.BusBookings))
	})
	t.Run("success::not_empty", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		const (
			route = "ufaspb"
			date  = "2020-01-01"
			seat  = uint32(1)
		)
		ctx := context.Background()
		requestAdd := &pb.BusBookingAddRequest{
			Route: route,
			Date:  date,
			Seat:  seat,
		}
		request := pb.BusBookingListRequest{
			Offset: 0,
			Limit:  100,
		}

		// Act
		responseAdd, errAdd := BusBookingClient.BusBookingAdd(ctx, requestAdd)
		require.NoError(t, errAdd)
		response, err := BusBookingClient.BusBookingList(ctx, &request)

		// Assert
		require.NoError(t, err)
		require.Equal(t, 1, len(response.BusBookings))
		assert.Equal(t, &pb.BusBooking{
			Id:    responseAdd.GetId(),
			Route: route,
			Date:  date,
			Seat:  seat,
		}, response.GetBusBookings()[0])
	})
}

func TestAddBusBookingIntegration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: "ufaspb",
			Date:  "2020-01-01",
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
			Date:  "2020-01-01",
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
	t.Run("fail::incorrect_date", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: "ufaspb",
			Date:  "TEST",
			Seat:  1,
		}

		// Act
		_, err := BusBookingClient.BusBookingAdd(ctx, &request)
		st, ok := status.FromError(err)

		// Assert
		require.True(t, ok)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Contains(t, st.Message(), "expected correct date in format")
	})
	t.Run("fail::incorrect_seat", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: "ufaspb",
			Date:  "2020-01-01",
			Seat:  101,
		}

		// Act
		_, err := BusBookingClient.BusBookingAdd(ctx, &request)
		st, ok := status.FromError(err)

		// Assert
		require.True(t, ok)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Contains(t, st.Message(), "expected seat number from")
	})
	t.Run("fail::route_not_exist", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: "abcdef",
			Date:  "2020-01-01",
			Seat:  1,
		}

		// Act
		_, err := BusBookingClient.BusBookingAdd(ctx, &request)
		st, ok := status.FromError(err)

		// Assert
		require.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Equal(t, st.Message(), "route not exist")
	})
	t.Run("fail::bb_already_exist", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingAddRequest{
			Route: "ufaspb",
			Date:  "2020-01-01",
			Seat:  1,
		}
		// Act
		_, err1 := BusBookingClient.BusBookingAdd(ctx, &request)
		_, err2 := BusBookingClient.BusBookingAdd(ctx, &request)
		st, ok := status.FromError(err2)

		// Assert
		require.NoError(t, err1)
		require.True(t, ok)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Equal(t, st.Message(), "bus booking already exists")
	})
}

func TestGetBusBookingIntegration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		const (
			route = "ufaspb"
			date  = "2020-01-01"
			seat  = uint32(1)
		)
		ctx := context.Background()
		requestAdd := &pb.BusBookingAddRequest{
			Route: route,
			Date:  date,
			Seat:  seat,
		}

		// Act
		responseAdd, errAdd := BusBookingClient.BusBookingAdd(ctx, requestAdd)
		require.NoError(t, errAdd)
		response, err := BusBookingClient.BusBookingGet(ctx, &pb.BusBookingGetRequest{
			Id: responseAdd.GetId(),
		})

		// Assert
		require.NoError(t, err)
		assert.Equal(t, &pb.BusBooking{
			Id:    responseAdd.GetId(),
			Route: route,
			Date:  date,
			Seat:  seat,
		}, response.GetBusBooking())
	})
	t.Run("fail::no_exist", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingGetRequest{
			Id: uint32(1),
		}

		// Act
		_, err := BusBookingClient.BusBookingGet(ctx, &request)
		st, ok := status.FromError(err)

		// Assert
		require.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Equal(t, st.Message(), "bus booking does not exist")
	})
}

func TestChangeSeatBusBookingIntegration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		const (
			route   = "ufaspb"
			date    = "2020-01-01"
			seat    = uint32(1)
			newSeat = uint32(2)
		)
		ctx := context.Background()
		requestAdd := &pb.BusBookingAddRequest{
			Route: route,
			Date:  date,
			Seat:  seat,
		}

		// Act
		responseAdd, errAdd := BusBookingClient.BusBookingAdd(ctx, requestAdd)
		require.NoError(t, errAdd)
		_, err := BusBookingClient.BusBookingChangeSeat(ctx, &pb.BusBookingChangeSeatRequest{
			Id:   responseAdd.GetId(),
			Seat: newSeat,
		})

		// Assert
		assert.NoError(t, err)
	})
	t.Run("fail::no_exist", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingChangeSeatRequest{
			Id:   uint32(1),
			Seat: uint32(2),
		}

		// Act
		_, err := BusBookingClient.BusBookingChangeSeat(ctx, &request)
		st, ok := status.FromError(err)

		// Assert
		require.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Equal(t, st.Message(), "bus booking does not exist")
	})
}

func TestChangeDateSeatBusBookingIntegration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		const (
			route   = "ufaspb"
			date    = "2020-01-01"
			newDate = "2020-01-02"
			seat    = uint32(1)
			newSeat = uint32(2)
		)
		ctx := context.Background()
		requestAdd := &pb.BusBookingAddRequest{
			Route: route,
			Date:  date,
			Seat:  seat,
		}

		// Act
		responseAdd, errAdd := BusBookingClient.BusBookingAdd(ctx, requestAdd)
		require.NoError(t, errAdd)
		_, err := BusBookingClient.BusBookingChangeDateSeat(ctx, &pb.BusBookingChangeDateSeatRequest{
			Id:   responseAdd.GetId(),
			Date: newDate,
			Seat: newSeat,
		})

		// Assert
		assert.NoError(t, err)
	})
	t.Run("fail::no_exist", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingChangeDateSeatRequest{
			Id:   uint32(1),
			Date: "2020-01-01",
			Seat: uint32(2),
		}

		// Act
		_, err := BusBookingClient.BusBookingChangeDateSeat(ctx, &request)
		st, ok := status.FromError(err)

		// Assert
		require.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Equal(t, st.Message(), "bus booking does not exist")
	})
}

func TestDeleteBusBookingIntegration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		requestAdd := &pb.BusBookingAddRequest{
			Route: "ufaspb",
			Date:  "2020-01-01",
			Seat:  uint32(1),
		}

		// Act
		responseAdd, errAdd := BusBookingClient.BusBookingAdd(ctx, requestAdd)
		require.NoError(t, errAdd)
		_, err := BusBookingClient.BusBookingDelete(ctx, &pb.BusBookingDeleteRequest{
			Id: responseAdd.GetId(),
		})

		// Assert
		assert.NoError(t, err)
	})
	t.Run("fail::no_exist", func(t *testing.T) {
		// Arrange
		TestDatabase.Setup(t)
		defer TestDatabase.TearDown(t)

		ctx := context.Background()
		request := pb.BusBookingDeleteRequest{
			Id: uint32(1),
		}

		// Act
		_, err := BusBookingClient.BusBookingDelete(ctx, &request)
		st, ok := status.FromError(err)

		// Assert
		require.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Equal(t, st.Message(), "bus booking does not exist")
	})
}
