package get

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
	return "get"
}

func (c *command) Description() string {
	return "get bus booking by id"
}

func (c *command) Process(args string) string {
	params, err := commandPkg.CheckArguments(args, 1)
	if err != nil {
		return err.Error()
	}

	id, err := toolsPkg.StringToUint(params[0])
	if err != nil {
		return err.Error()
	}

	if bb, err := c.bb.Get(id); err != nil {
		return err.Error()
	} else {
		return fmt.Sprintf("[SUCCESS]\n%d: %s / %s / %d", id, bb.Route, bb.Date, bb.Seat)
	}
}
