package change_seat

import (
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
	return "change_seat"
}

func (c *command) Description() string {
	return "change seat of bus booking"
}

func (c *command) Process(args string) string {
	params, err := commandPkg.CheckArguments(args, 2)
	if err != nil {
		return err.Error()
	}

	id, err := toolsPkg.StringToUint(params[0])
	if err != nil {
		return err.Error()
	}
	newSeat, err := toolsPkg.StringToUint(params[1])
	if err != nil {
		return err.Error()
	}

	if err := c.bb.ChangeSeat(id, newSeat); err != nil {
		return err.Error()
	}
	return fmt.Sprintf("[SUCCESS]\nBooking seat changed successfully to %d\nbooking number: %d", newSeat, id)
}
