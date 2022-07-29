package main

import (
	"log"

	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/commander"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/handlers"
)

func main() {
	log.Println("start main")
	cmd := commander.MustNew()
	handlers.AddHandlers(cmd)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
