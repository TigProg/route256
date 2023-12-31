package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

type PoolInterface interface {
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Close()
}

func New(pool PoolInterface) repoPkg.Interface {
	return &repo{pool}
}

type repo struct {
	pool PoolInterface
}

func (r *repo) List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error) {
	query := `
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
	`
	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		log.Errorf("postgresRepoPkg::List query error %s", err.Error())
		return nil, repoPkg.ErrRepoInternal
	}
	defer rows.Close()

	var result []models.BusBooking
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Errorf("postgresRepoPkg::List empty rows after check %s", err.Error())
			return nil, repoPkg.ErrRepoInternal
		}
		result = append(result, models.BusBooking{
			Id:    uint(values[0].(int32)),
			Route: values[1].(string),
			Date:  values[2].(string),
			Seat:  uint(values[3].(int32)),
		})
	}
	return result, nil
}

func (r *repo) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	existedRouteId, err := r.getRouteIdByRouteName(ctx, bb.Route)
	if err != nil {
		if errors.Is(err, repoPkg.ErrRepoRouteNameNotExist) {
			return 0, err
		}
		log.Errorf("postgresRepoPkg::Add getRouteIdByRouteName %s", err.Error())
		return 0, repoPkg.ErrRepoInternal
	}

	_, err = r.reverseSearch(ctx, bb.Route, bb.Date, bb.Seat)
	if err == nil {
		return 0, repoPkg.ErrRepoBusBookingAlreadyExists
	}
	if !errors.Is(err, repoPkg.ErrRepoBusBookingNotExists) {
		log.Errorf("postgresRepoPkg::Add reverseSearch %s", err.Error())
		return 0, repoPkg.ErrRepoInternal
	}

	query := `
		INSERT INTO public.booking (route_id, date, seat) VALUES
		($1, $2, $3)
		RETURNING id
	`
	rows, err := r.pool.Query(ctx, query, existedRouteId, bb.Date, bb.Seat)
	if err != nil {
		log.Errorf("postgresRepoPkg::Add query error %s", err.Error())
		return 0, repoPkg.ErrRepoInternal
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Errorf("postgresRepoPkg::Add empty rows after check %s", err.Error())
			return 0, repoPkg.ErrRepoInternal
		}
		return uint(values[0].(int32)), nil
	}
	log.Errorf("postgresRepoPkg::Add empty returning %s", err.Error())
	return 0, repoPkg.ErrRepoInternal
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
		log.Errorf("postgresRepoPkg::Get query error %s", err.Error())
		return nil, repoPkg.ErrRepoInternal
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Errorf("postgresRepoPkg::Get empty rows after check %s", err.Error())
			return nil, repoPkg.ErrRepoInternal
		}
		return &models.BusBooking{
			Id:    uint(values[0].(int32)),
			Route: values[1].(string),
			Date:  values[2].(string),
			Seat:  uint(values[3].(int32)),
		}, nil
	}
	return nil, repoPkg.ErrRepoBusBookingNotExists
}

func (r *repo) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	existedBb, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	if existedBb.Seat == newSeat {
		return nil // for idempotency
	}

	_, err = r.reverseSearch(ctx, existedBb.Route, existedBb.Date, newSeat)
	if err == nil {
		return repoPkg.ErrRepoBusBookingAlreadyExists
	}
	if !errors.Is(err, repoPkg.ErrRepoBusBookingNotExists) {
		log.Errorf("postgresRepoPkg::ChangeSeat reverseSearch %s", err.Error())
		return repoPkg.ErrRepoInternal
	}

	query := `
		UPDATE public.booking
		SET seat = $2
		WHERE id = $1
	`
	rows, err := r.pool.Query(ctx, query, id, newSeat)
	if err != nil {
		log.Errorf("postgresRepoPkg::ChangeSeat query error %s", err.Error())
		return repoPkg.ErrRepoInternal
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

	_, err = r.reverseSearch(ctx, existedBb.Route, newDate, newSeat)
	if err == nil {
		return repoPkg.ErrRepoBusBookingAlreadyExists
	}
	if !errors.Is(err, repoPkg.ErrRepoBusBookingNotExists) {
		log.Errorf("postgresRepoPkg::ChangeDateSeat reverseSearch %s", err.Error())
		return repoPkg.ErrRepoInternal
	}

	query := `
		UPDATE public.booking
		SET seat = $2, date = $3
		WHERE id = $1
	`
	rows, err := r.pool.Query(ctx, query, id, newSeat, newDate)
	if err != nil {
		log.Errorf("postgresRepoPkg::ChangeDateSeat query error %s", err.Error())
		return repoPkg.ErrRepoInternal
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
		FROM public.booking
		WHERE id = $1
	`
	rows, err := r.pool.Query(ctx, query, id)
	if err != nil {
		log.Errorf("postgresRepoPkg::Delete query error %s", err.Error())
		return repoPkg.ErrRepoInternal
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
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return 0, err
		}
		return uint(values[0].(int32)), nil
	}
	return 0, repoPkg.ErrRepoBusBookingNotExists
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
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return 0, err
		}
		return uint(values[0].(int32)), nil
	}
	return 0, repoPkg.ErrRepoRouteNameNotExist
}
