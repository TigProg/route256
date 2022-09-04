package redis_repo

import (
	"context"

	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
	metricPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/metrics"
	redisWrapperPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/redis_wrapper"
)

func New(readRepo, writeRepo repoPkg.Interface, client redisWrapperPkg.Interface) repoPkg.InterfaceWithMetrics {
	return &repo{
		readRepo:  readRepo,
		writeRepo: writeRepo,
		client:    client,
		metrics: metricManager{
			hit:  metricPkg.NewMetric("redis_repo::hit"),
			miss: metricPkg.NewMetric("redis_repo::miss"),
		},
	}
}

// TODO change structure to remove copy-paste
type metricManager struct {
	hit  *metricPkg.Metric
	miss *metricPkg.Metric
}

type repo struct {
	readRepo  repoPkg.Interface
	writeRepo repoPkg.Interface
	client    redisWrapperPkg.Interface
	metrics   metricManager
}

func (r *repo) GetMetrics() []*metricPkg.Metric {
	return []*metricPkg.Metric{
		r.metrics.hit,
		r.metrics.miss,
	}
}

func (r *repo) List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error) {
	return r.readRepo.List(ctx, offset, limit)
}

func (r *repo) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	id, err := r.writeRepo.Add(ctx, bb)
	if err != nil {
		err := r.client.SetById(id, bb)
		if err != nil {
			_ = r.client.DisableById(id)
		}
	}
	return id, err
}

func (r *repo) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	bb, err := r.client.GetById(id)
	if err == nil {
		r.metrics.hit.Increment()
		return bb, err
	}
	r.metrics.miss.Increment()

	bb, err = r.readRepo.Get(ctx, id)
	if err == nil {
		err = r.client.SetById(id, *bb)
		if err != nil {
			_ = r.client.DisableById(id)
		}
	}
	return bb, err
}

func (r *repo) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	_ = r.client.DisableById(id)
	return r.writeRepo.ChangeSeat(ctx, id, newSeat)
}

func (r *repo) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	_ = r.client.DisableById(id)
	return r.writeRepo.ChangeDateSeat(ctx, id, newDate, newSeat)
}

func (r *repo) Delete(ctx context.Context, id uint) error {
	_ = r.client.DisableById(id)
	return r.writeRepo.Delete(ctx, id)
}
