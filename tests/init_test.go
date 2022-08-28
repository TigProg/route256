//go:build integration
// +build integration

package tests

import (
	"log"
	"time"

	"gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
	"gitlab.ozon.dev/tigprog/bus_booking/tests/config"
	"gitlab.ozon.dev/tigprog/bus_booking/tests/postgres"

	"google.golang.org/grpc"
)

var (
	BusBookingClient api.AdminClient
	TestDatabase     *postgres.TDB
)

func init() {
	cfg, err := config.FromEnv()
	conn, err := grpc.Dial(cfg.Host, grpc.WithInsecure(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Panic(err)
	}
	BusBookingClient = api.NewAdminClient(conn)
	TestDatabase = postgres.New(cfg)
}
