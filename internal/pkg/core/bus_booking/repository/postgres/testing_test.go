package postgres

import (
	"testing"

	"github.com/pashagolub/pgxmock"
	repoPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/repository"
)

type busBookingRepoFixtures struct {
	busBookingRepo repoPkg.Interface
	dbPoolMock     pgxmock.PgxPoolIface
}

func setUp(t *testing.T) busBookingRepoFixtures {
	var fixture busBookingRepoFixtures

	poolMock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error") // TODO
	}
	fixture.dbPoolMock = poolMock
	fixture.busBookingRepo = New(fixture.dbPoolMock)
	return fixture
}

func (f *busBookingRepoFixtures) tearDown() {
	f.dbPoolMock.Close()
}
