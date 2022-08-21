package api

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	pb "gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
)

func TestListBusBookingApi(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)

		const (
			offset = uint(0)
			limit  = uint(100)
		)

		request := pb.BusBookingListRequest{
			Offset: uint32(offset),
			Limit:  uint32(limit),
		}
		bbs := []models.BusBooking{
			{
				Id:    uint(1),
				Route: "spbufa",
				Date:  "2000-01-01",
				Seat:  uint(10),
			},
			{
				Id:    uint(2),
				Route: "ufamsk",
				Date:  "2001-01-01",
				Seat:  uint(15),
			},
		}
		f.busBookingRepo.EXPECT().
			List(gomock.Any(), offset, limit).
			Return(bbs, nil)

		// Act
		response, err := f.service.BusBookingList(f.ctx, &request)
		expected := &pb.BusBookingListResponse{
			BusBookings: []*pb.BusBooking{
				{
					Id:    uint32(1),
					Route: "spbufa",
					Date:  "2000-01-01",
					Seat:  uint32(10),
				},
				{
					Id:    uint32(2),
					Route: "ufamsk",
					Date:  "2001-01-01",
					Seat:  uint32(15),
				},
			},
		}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expected, response)
	})
}

func TestAddBusBookingApi(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)

		const (
			bbId    = uint(1)
			bbRoute = "ufaspb"
			bbDate  = "2020-01-01"
			bbSeat  = uint(2)
		)

		request := pb.BusBookingAddRequest{
			Route: bbRoute,
			Date:  bbDate,
			Seat:  uint32(bbSeat),
		}
		f.busBookingRepo.EXPECT().
			Add(gomock.Any(), models.BusBooking{
				Id:    0,
				Route: bbRoute,
				Date:  bbDate,
				Seat:  bbSeat,
			}).
			Return(bbId, nil)

		// Act
		response, err := f.service.BusBookingAdd(f.ctx, &request)
		expected := &pb.BusBookingAddResponse{Id: uint32(bbId)}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expected, response)
	})
}

func TestGetBusBookingApi(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)

		const (
			bbId    = uint(1)
			bbRoute = "ufaspb"
			bbDate  = "2020-01-01"
			bbSeat  = uint(2)
		)

		request := pb.BusBookingGetRequest{Id: uint32(bbId)}
		bb := &models.BusBooking{
			Id:    bbId,
			Route: bbRoute,
			Date:  bbDate,
			Seat:  bbSeat,
		}
		f.busBookingRepo.EXPECT().
			Get(gomock.Any(), bbId).
			Return(bb, nil)

		// Act
		response, err := f.service.BusBookingGet(f.ctx, &request)
		expected := &pb.BusBookingGetResponse{
			BusBooking: &pb.BusBooking{
				Id:    uint32(bbId),
				Route: bbRoute,
				Date:  bbDate,
				Seat:  uint32(bbSeat),
			},
		}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expected, response)
	})
}

func TestChangeSeatBusBookingApi(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)

		const (
			bbId      = uint(1)
			bbNewSeat = uint(3)
		)

		request := pb.BusBookingChangeSeatRequest{
			Id:   uint32(bbId),
			Seat: uint32(bbNewSeat),
		}
		f.busBookingRepo.EXPECT().
			ChangeSeat(gomock.Any(), bbId, bbNewSeat).
			Return(nil)

		// Act
		response, err := f.service.BusBookingChangeSeat(f.ctx, &request)
		expected := &pb.BusBookingChangeSeatResponse{}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expected, response)
	})
}

func TestChangeDateSeatBusBookingApi(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)

		const (
			bbId      = uint(1)
			bbNewDate = "2020-01-01"
			bbNewSeat = uint(3)
		)

		request := pb.BusBookingChangeDateSeatRequest{
			Id:   uint32(bbId),
			Date: bbNewDate,
			Seat: uint32(bbNewSeat),
		}
		f.busBookingRepo.EXPECT().
			ChangeDateSeat(gomock.Any(), bbId, bbNewDate, bbNewSeat).
			Return(nil)

		// Act
		response, err := f.service.BusBookingChangeDateSeat(f.ctx, &request)
		expected := &pb.BusBookingChangeDateSeatResponse{}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expected, response)
	})
}

func TestDeleteBusBookingApi(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)

		const bbId = uint(1)

		request := pb.BusBookingDeleteRequest{
			Id: uint32(bbId),
		}
		f.busBookingRepo.EXPECT().
			Delete(gomock.Any(), bbId).
			Return(nil)

		// Act
		response, err := f.service.BusBookingDelete(f.ctx, &request)
		expected := &pb.BusBookingDeleteResponse{}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expected, response)
	})
}
