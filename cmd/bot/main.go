package main

import bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"

func main() {
	bb := bbPkg.New()

	go runBot(bb)
	go runREST()
	runGRPCServer(bb)
}
