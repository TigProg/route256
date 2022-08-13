package add

import (
	"context"
	"fmt"

	commandPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	"gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking/models"
)

func New(bb bbPkg.Interface) commandPkg.Interface {
	return &command{
		bb: bb,
	}
}

type command struct {
	bb bbPkg.Interface
}

func (c *command) Name() string {
	return "add"
}

func (c *command) Description() string {
	return "create bus booking"
}

func (c *command) Process(ctx context.Context, args string) string {
	params, err := commandPkg.CheckArguments(args, 3)
	if err != nil {
		return err.Error()
	}

	seat, err := commandPkg.StringToUint(params[2])
	if err != nil {
		return err.Error()
	}

	if id, err := c.bb.Add(ctx, models.BusBooking{
		Id:    0,
		Route: params[0],
		Date:  params[1],
		Seat:  seat,
	}); err != nil {
		return err.Error()
	} else {
		return fmt.Sprintf("[SUCCESS]\nBus seat booked successfully\nbooking number: %d", id)
	}
}
