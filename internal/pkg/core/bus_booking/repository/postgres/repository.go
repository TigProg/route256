package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

var (
	ErrInternal          = errors.New("internal error")
	ErrRouteNameNotExist = errors.New("route not exist")
)

func New(pool *pgxpool.Pool) repoPkg.Interface {
	return &repo{pool}
}

type repo struct {
	pool *pgxpool.Pool
}

func (r *repo) List(ctx context.Context) ([]models.BusBooking, error) {
	query := `
		SELECT
			b.id AS id,
			r.name AS route,
			b.date::varchar(10) AS date,
			b.seat AS seat
		FROM
			public.booking as b
			INNER JOIN public.route as r ON (b.route_id = r.id)
	`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		// TODO log
		return nil, ErrInternal
	}
	defer rows.Close()

	var result []models.BusBooking
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			// TODO log
			return nil, ErrInternal
		}
		result = append(result, models.BusBooking{
			Id:    values[0].(uint),
			Route: values[1].(string),
			Date:  values[2].(string),
			Seat:  values[3].(uint),
		})
	}
	return result, nil
}

func (r *repo) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	existedId, err := r.reverseSearch(ctx, bb.Route, bb.Date, bb.Seat)
	if err == nil {
		return 0, errors.Wrapf(repoPkg.ErrBusBookingAlreadyExists, "%d", existedId)
	}
	if !errors.Is(err, repoPkg.ErrBusBookingNotExists) {
		// TODO log
		return 0, ErrInternal
	}

	existedRouteId, err := r.getRouteIdByRouteName(ctx, bb.Route)
	if err != nil {
		if errors.Is(err, ErrRouteNameNotExist) {
			return 0, err
		}
		return 0, ErrInternal
	}

	query := `
		INSERT INTO public.booking (route_id, date, seat) VALUES
		($1, $2, $3)
		RETURNING id
	`
	rows, err := r.pool.Query(ctx, query, existedRouteId, bb.Date, bb.Seat)
	if err != nil {
		// TODO log
		return 0, ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			// TODO log
			return 0, ErrInternal
		}
		return values[0].(uint), nil
	}
	return 0, ErrInternal
}

func (r *repo) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	query := `
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
	`
	rows, err := r.pool.Query(ctx, query, id)
	if err != nil {
		// TODO log
		return nil, ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			// TODO log
			return nil, ErrInternal
		}
		return &models.BusBooking{
			Id:    values[0].(uint),
			Route: values[1].(string),
			Date:  values[2].(string),
			Seat:  values[3].(uint),
		}, nil
	}
	return nil, repoPkg.ErrBusBookingNotExists
}

func (r *repo) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	existedBb, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	if existedBb.Seat == newSeat {
		return nil // for idempotency
	}

	existedId, err := r.reverseSearch(ctx, existedBb.Route, existedBb.Date, newSeat)
	if err == nil {
		return errors.Wrapf(repoPkg.ErrBusBookingAlreadyExists, "%d", existedId)
	}
	if !errors.Is(err, repoPkg.ErrBusBookingNotExists) {
		// TODO log
		return ErrInternal
	}

	query := `
		UPDATE public.booking
		SET seat = $2
		WHERE id = $1
	`
	rows, err := r.pool.Query(ctx, query, id, newSeat)
	if err != nil {
		// TODO log
		return ErrInternal
	}
	defer rows.Close()
	return nil
}

func (r *repo) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	existedBb, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	if existedBb.Date == newDate && existedBb.Seat == newSeat {
		return nil // for idempotency
	}

	existedId, err := r.reverseSearch(ctx, existedBb.Route, newDate, newSeat)
	if err == nil {
		return errors.Wrapf(repoPkg.ErrBusBookingAlreadyExists, "%d", existedId)
	}
	if !errors.Is(err, repoPkg.ErrBusBookingNotExists) {
		// TODO log
		return ErrInternal
	}

	query := `
		UPDATE public.booking
		SET seat = $2, date = $3
		WHERE id = $1
	`
	rows, err := r.pool.Query(ctx, query, id, newSeat, newDate)
	if err != nil {
		// TODO log
		return ErrInternal
	}
	defer rows.Close()
	return nil
}

func (r *repo) Delete(ctx context.Context, id uint) error {
	_, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	query := `
		DELETE
		FROM public.route
		WHERE id = $1
	`
	rows, err := r.pool.Query(ctx, query, id)
	if err != nil {
		// TODO log
		return ErrInternal
	}
	defer rows.Close()
	return nil
}

func (r *repo) reverseSearch(ctx context.Context, route string, date string, seat uint) (uint, error) {
	query := `
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
	`
	rows, err := r.pool.Query(ctx, query, route, date, seat)
	if err != nil {
		// TODO log
		return 0, ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			// TODO log
			return 0, ErrInternal
		}
		return values[0].(uint), nil
	}
	return 0, repoPkg.ErrBusBookingNotExists
}

func (r *repo) getRouteIdByRouteName(ctx context.Context, routeName string) (uint, error) {
	query := `
		SELECT
			id
		FROM
			public.route
		WHERE
		    name = $1
	`
	rows, err := r.pool.Query(ctx, query, routeName)
	if err != nil {
		// TODO log
		return 0, ErrInternal
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			// TODO log
			return 0, ErrInternal
		}
		return values[0].(uint), nil
	}
	return 0, ErrRouteNameNotExist
}
