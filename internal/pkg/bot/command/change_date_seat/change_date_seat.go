package change_date_seat

import (
	"context"
	"fmt"

	commandPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
	toolsPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/tools"
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
	return "change_date_seat"
}

func (c *command) Description() string {
	return "change date and seat of bus booking"
}

func (c *command) Process(ctx context.Context, args string) string {
	params, err := commandPkg.CheckArguments(args, 3)
	if err != nil {
		return err.Error()
	}

	id, err := toolsPkg.StringToUint(params[0])
	if err != nil {
		return err.Error()
	}
	newSeat, err := toolsPkg.StringToUint(params[2])
	if err != nil {
		return err.Error()
	}

	if err := c.bb.ChangeDateSeat(ctx, id, params[1], newSeat); err != nil {
		return err.Error()
	}
	return fmt.Sprintf(
		"[SUCCESS]\nBooking date and seat changed successfully to %s, %d\nbooking number: %d",
		params[1], newSeat, id,
	)
}
