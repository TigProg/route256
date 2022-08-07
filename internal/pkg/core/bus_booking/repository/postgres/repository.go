package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

var ErrSome = errors.New("some error")

func New(pool *pgxpool.Pool) repoPkg.Interface {
	return &repo{pool}
}

type repo struct {
	pool *pgxpool.Pool
}

func (r *repo) List(ctx context.Context) []models.BusBooking {
	var bbs []models.BusBooking
	return bbs // TODO
}

func (r *repo) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	return 0, ErrSome // TODO
}

func (r *repo) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	query := `
		SELECT
			b.id AS id,
			r.name AS route,
			b.date::varchar(10) AS date,
			b.seat AS seat
		FROM
			public.booking as b
			INNER JOIN public.route as r ON (b.route_id = r.id)
		WHERE
			b.id = $1
	`
	rows, err := r.pool.Query(ctx, query, id)
	if err != nil {
		// TODO log
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			// TODO log
			return nil, err
		}
		return &models.BusBooking{
			Id:    values[0].(uint),
			Route: values[1].(string),
			Date:  values[2].(string),
			Seat:  values[3].(uint),
		}, err
	}
	return nil, repoPkg.ErrBusBookingNotExists
}

func (r *repo) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	return ErrSome // TODO
}

func (r *repo) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	return ErrSome // TODO
}

func (r *repo) Delete(ctx context.Context, id uint) error {
	return ErrSome // TODO
}
