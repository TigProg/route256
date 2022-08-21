//go:build integration
// +build integration

package tests

import (
	"log"
	"time"

	"gitlab.ozon.dev/tigprog/bus_booking/pkg/api"
	"gitlab.ozon.dev/tigprog/bus_booking/tests/config"

	"google.golang.org/grpc"
)

var (
	BusBookingClient api.AdminClient
)

func init() {
	cfg, err := config.FromEnv()
	conn, err := grpc.Dial(cfg.Host, grpc.WithInsecure(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Panic(err)
	}
	BusBookingClient = api.NewAdminClient(conn)
}
