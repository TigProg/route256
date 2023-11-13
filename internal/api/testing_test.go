package api

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	mockRepository "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository/mocks"
)

type busBookingApiFixture struct {
	ctx            context.Context
	busBookingRepo *mockRepository.MockInterface
	service        *Implementation
}

func setUp(t *testing.T) busBookingApiFixture {
	t.Parallel() // TODO

	ctx := context.Background()

	f := busBookingApiFixture{}
	f.ctx = ctx
	f.busBookingRepo = mockRepository.NewMockInterface(gomock.NewController(t))
	f.service = New(bbPkg.New(f.busBookingRepo))
	return f
}
