package rw_repo

import (
	"context"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

func New(readRepo, writeRepo repoPkg.Interface) repoPkg.Interface {
	return &repo{
		readRepo:  readRepo,
		writeRepo: writeRepo,
	}
}

type repo struct {
	readRepo  repoPkg.Interface
	writeRepo repoPkg.Interface
}

func (r *repo) List(ctx context.Context, offset uint, limit uint) ([]models.BusBooking, error) {
	return r.readRepo.List(ctx, offset, limit)
}

func (r *repo) Add(ctx context.Context, bb models.BusBooking) (uint, error) {
	return r.writeRepo.Add(ctx, bb)
}

func (r *repo) Get(ctx context.Context, id uint) (*models.BusBooking, error) {
	return r.readRepo.Get(ctx, id)
}

func (r *repo) ChangeSeat(ctx context.Context, id uint, newSeat uint) error {
	return r.writeRepo.ChangeSeat(ctx, id, newSeat)
}

func (r *repo) ChangeDateSeat(ctx context.Context, id uint, newDate string, newSeat uint) error {
	return r.writeRepo.ChangeDateSeat(ctx, id, newDate, newSeat)
}

func (r *repo) Delete(ctx context.Context, id uint) error {
	return r.writeRepo.Delete(ctx, id)
}
