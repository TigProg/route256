package postgres

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

func TestListBusBookingPostgresRepo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"}).
			AddRow(int32(1), "spbufa", "2000-01-01", int32(10)).
			AddRow(int32(2), "ufamsk", "2001-01-01", int32(15)).
			AddRow(int32(3), "mskspb", "2002-01-01", int32(20))
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking as b
				INNER JOIN public.route as r ON (b.route_id = r.id)
			ORDER BY
				b.id
			LIMIT
				$1
			OFFSET
				$2
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(uint(100), uint(0)).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		result, err := f.busBookingRepo.List(ctx, 0, 100)
		expected := []models.BusBooking{
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
			{
				Id:    uint(3),
				Route: "mskspb",
				Date:  "2002-01-01",
				Seat:  uint(20),
			},
		}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("fail::db_not_available", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking as b
				INNER JOIN public.route as r ON (b.route_id = r.id)
			ORDER BY
				b.id
			LIMIT
				$1
			OFFSET
				$2
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(uint(100), uint(0)).
			WillReturnError(errors.New("some db error"))

		// Act
		ctx := context.Background()
		_, err := f.busBookingRepo.List(ctx, 0, 100)

		// Assert
		assert.Equal(t, repoPkg.ErrRepoInternal, err)
	})
}

func TestAddBusBookingPostgresRepo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId     = uint(0)
			bbRealId = uint(1)
			bbRoute  = "spbufa"
			routeId  = uint(2)
			bbDate   = "2020-01-01"
			bbSeat   = uint(3)
		)

		// getRouteIdByRouteName
		rows := pgxmock.NewRows([]string{"id"}).AddRow(int32(routeId))
		queryStore := regexp.QuoteMeta(`
			SELECT
				id
			FROM
				public.route
			WHERE
				name = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute).
			WillReturnRows(rows)

		// reverseSearch
		queryStore = regexp.QuoteMeta(`
			SELECT
				b.id AS id
			FROM
				public.booking AS b
				INNER JOIN (
					SELECT
						id
					FROM
						public.route
					WHERE
						name = $1
				) AS r ON (b.route_id = r.id)
			WHERE
				b.date = $2
				AND b.seat = $3
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute, bbDate, bbSeat).
			WillReturnError(repoPkg.ErrRepoBusBookingNotExists)

		// insert
		rows = pgxmock.NewRows([]string{"id"}).AddRow(int32(bbRealId))
		queryStore = regexp.QuoteMeta(`
			INSERT INTO public.booking (route_id, date, seat) VALUES
			($1, $2, $3)
			RETURNING id
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(routeId, bbDate, bbSeat).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		result, err := f.busBookingRepo.Add(ctx, models.BusBooking{
			Id:    bbId,
			Route: bbRoute,
			Date:  bbDate,
			Seat:  bbSeat,
		})

		// Assert
		require.NoError(t, err)
		assert.Equal(t, bbRealId, result)
	})
	t.Run("fail::route_name_not_exist", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId    = uint(0)
			bbRoute = "spbufa"
			bbDate  = "2020-01-01"
			bbSeat  = uint(3)
		)

		// getRouteIdByRouteName
		rows := pgxmock.NewRows([]string{"id"})
		queryStore := regexp.QuoteMeta(`
			SELECT
				id
			FROM
				public.route
			WHERE
				name = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		_, err := f.busBookingRepo.Add(ctx, models.BusBooking{
			Id:    bbId,
			Route: bbRoute,
			Date:  bbDate,
			Seat:  bbSeat,
		})

		// Assert
		assert.Equal(t, repoPkg.ErrRepoRouteNameNotExist, err)
	})
	t.Run("fail::bus_booking_already_exist", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId        = uint(0)
			bbExistedId = uint(1)
			bbRoute     = "spbufa"
			routeId     = uint(2)
			bbDate      = "2020-01-01"
			bbSeat      = uint(3)
		)

		// getRouteIdByRouteName
		rows := pgxmock.NewRows([]string{"id"}).AddRow(int32(routeId))
		queryStore := regexp.QuoteMeta(`
			SELECT
				id
			FROM
				public.route
			WHERE
				name = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute).
			WillReturnRows(rows)

		// reverseSearch
		rows = pgxmock.NewRows([]string{"id"}).AddRow(int32(bbExistedId))
		queryStore = regexp.QuoteMeta(`
			SELECT
				b.id AS id
			FROM
				public.booking AS b
				INNER JOIN (
					SELECT
						id
					FROM
						public.route
					WHERE
						name = $1
				) AS r ON (b.route_id = r.id)
			WHERE
				b.date = $2
				AND b.seat = $3
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute, bbDate, bbSeat).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		_, err := f.busBookingRepo.Add(ctx, models.BusBooking{
			Id:    bbId,
			Route: bbRoute,
			Date:  bbDate,
			Seat:  bbSeat,
		})

		// Assert
		assert.Equal(t, repoPkg.ErrRepoBusBookingAlreadyExists, err)
	})
}

func TestGetBusBookingPostgresRepo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId    = uint(0)
			bbRoute = "spbufa"
			bbDate  = "2020-01-01"
			bbSeat  = uint(1)
		)

		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"}).
			AddRow(int32(bbId), bbRoute, bbDate, int32(bbSeat))
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking AS b
				INNER JOIN public.route AS r ON (b.route_id = r.id)
			WHERE
				b.id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		result, err := f.busBookingRepo.Get(ctx, bbId)
		expected := &models.BusBooking{
			Id:    bbId,
			Route: bbRoute,
			Date:  bbDate,
			Seat:  bbSeat,
		}

		// Assert
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("fail::bus_booking_not_exist", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const bbId = uint(0)

		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"})
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking AS b
				INNER JOIN public.route AS r ON (b.route_id = r.id)
			WHERE
				b.id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		_, err := f.busBookingRepo.Get(ctx, bbId)

		// Assert
		assert.Equal(t, repoPkg.ErrRepoBusBookingNotExists, err)
	})
}

func TestChangeSeatBusBookingPostgresRepo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId      = uint(0)
			bbRoute   = "spbufa"
			bbDate    = "2020-01-01"
			bbSeat    = uint(1)
			bbNewSeat = uint(2)
		)

		// get
		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"}).
			AddRow(int32(bbId), bbRoute, bbDate, int32(bbSeat))
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking AS b
				INNER JOIN public.route AS r ON (b.route_id = r.id)
			WHERE
				b.id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// reverseSearch
		queryStore = regexp.QuoteMeta(`
			SELECT
				b.id AS id
			FROM
				public.booking AS b
				INNER JOIN (
					SELECT
						id
					FROM
						public.route
					WHERE
						name = $1
				) AS r ON (b.route_id = r.id)
			WHERE
				b.date = $2
				AND b.seat = $3
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute, bbDate, bbNewSeat).
			WillReturnError(repoPkg.ErrRepoBusBookingNotExists)

		// update
		rows = pgxmock.NewRows([]string{})
		queryStore = regexp.QuoteMeta(`
			UPDATE public.booking
			SET seat = $2
			WHERE id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId, bbNewSeat).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		err := f.busBookingRepo.ChangeSeat(ctx, bbId, bbNewSeat)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("fail::bus_booking_already_exist", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId        = uint(0)
			bbExistedId = uint(3)
			bbRoute     = "spbufa"
			bbDate      = "2020-01-01"
			bbSeat      = uint(1)
			bbNewSeat   = uint(2)
		)

		// get
		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"}).
			AddRow(int32(bbId), bbRoute, bbDate, int32(bbSeat))
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking AS b
				INNER JOIN public.route AS r ON (b.route_id = r.id)
			WHERE
				b.id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// reverseSearch
		rows = pgxmock.NewRows([]string{"id"}).AddRow(int32(bbExistedId))
		queryStore = regexp.QuoteMeta(`
			SELECT
				b.id AS id
			FROM
				public.booking AS b
				INNER JOIN (
					SELECT
						id
					FROM
						public.route
					WHERE
						name = $1
				) AS r ON (b.route_id = r.id)
			WHERE
				b.date = $2
				AND b.seat = $3
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute, bbDate, bbNewSeat).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		err := f.busBookingRepo.ChangeSeat(ctx, bbId, bbNewSeat)

		// Assert
		assert.Equal(t, repoPkg.ErrRepoBusBookingAlreadyExists, err)
	})
}

func TestChangeDateSeatBusBookingPostgresRepo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId      = uint(0)
			bbRoute   = "spbufa"
			bbDate    = "2020-01-01"
			bbSeat    = uint(1)
			bbNewDate = "2020-01-02"
			bbNewSeat = uint(2)
		)

		// get
		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"}).
			AddRow(int32(bbId), bbRoute, bbDate, int32(bbSeat))
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking AS b
				INNER JOIN public.route AS r ON (b.route_id = r.id)
			WHERE
				b.id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// reverseSearch
		queryStore = regexp.QuoteMeta(`
			SELECT
				b.id AS id
			FROM
				public.booking AS b
				INNER JOIN (
					SELECT
						id
					FROM
						public.route
					WHERE
						name = $1
				) AS r ON (b.route_id = r.id)
			WHERE
				b.date = $2
				AND b.seat = $3
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute, bbNewDate, bbNewSeat).
			WillReturnError(repoPkg.ErrRepoBusBookingNotExists)

		// update
		rows = pgxmock.NewRows([]string{})
		queryStore = regexp.QuoteMeta(`
			UPDATE public.booking
			SET seat = $2, date = $3
			WHERE id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId, bbNewSeat, bbNewDate).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		err := f.busBookingRepo.ChangeDateSeat(ctx, bbId, bbNewDate, bbNewSeat)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("fail::bus_booking_already_exist", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId        = uint(0)
			bbExistedId = uint(3)
			bbRoute     = "spbufa"
			bbDate      = "2020-01-01"
			bbSeat      = uint(1)
			bbNewDate   = "2020-01-02"
			bbNewSeat   = uint(2)
		)

		// get
		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"}).
			AddRow(int32(bbId), bbRoute, bbDate, int32(bbSeat))
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking AS b
				INNER JOIN public.route AS r ON (b.route_id = r.id)
			WHERE
				b.id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// reverseSearch
		rows = pgxmock.NewRows([]string{"id"}).AddRow(int32(bbExistedId))
		queryStore = regexp.QuoteMeta(`
			SELECT
				b.id AS id
			FROM
				public.booking AS b
				INNER JOIN (
					SELECT
						id
					FROM
						public.route
					WHERE
						name = $1
				) AS r ON (b.route_id = r.id)
			WHERE
				b.date = $2
				AND b.seat = $3
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbRoute, bbNewDate, bbNewSeat).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		err := f.busBookingRepo.ChangeDateSeat(ctx, bbId, bbNewDate, bbNewSeat)

		// Assert
		assert.Equal(t, repoPkg.ErrRepoBusBookingAlreadyExists, err)
	})
}

func TestDeleteBusBookingPostgresRepo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const (
			bbId    = uint(0)
			bbRoute = "spbufa"
			bbDate  = "2020-01-01"
			bbSeat  = uint(1)
		)

		// get
		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"}).
			AddRow(int32(bbId), bbRoute, bbDate, int32(bbSeat))
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking AS b
				INNER JOIN public.route AS r ON (b.route_id = r.id)
			WHERE
				b.id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// delete
		rows = pgxmock.NewRows([]string{})
		queryStore = regexp.QuoteMeta(`
			DELETE
			FROM public.booking
			WHERE id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		err := f.busBookingRepo.Delete(ctx, bbId)

		// Assert
		assert.NoError(t, err)
	})
	t.Run("fail::bus_booking_not_exist", func(t *testing.T) {
		// Arrange
		f := setUp(t)
		defer f.tearDown()

		const bbId = uint(0)

		// get
		rows := pgxmock.NewRows([]string{"id", "route", "date", "seat"})
		queryStore := regexp.QuoteMeta(`
			SELECT
				b.id AS id,
				r.name AS route,
				b.date::varchar(10) AS date,
				b.seat AS seat
			FROM
				public.booking AS b
				INNER JOIN public.route AS r ON (b.route_id = r.id)
			WHERE
				b.id = $1
		`)
		f.dbPoolMock.ExpectQuery(queryStore).
			WithArgs(bbId).
			WillReturnRows(rows)

		// Act
		ctx := context.Background()
		err := f.busBookingRepo.Delete(ctx, bbId)

		// Assert
		assert.Equal(t, repoPkg.ErrRepoBusBookingNotExists, err)
	})
}
