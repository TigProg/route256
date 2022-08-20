package list

import (
	"context"
	"fmt"

	commandPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command"
	bbPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/core/bus_booking"
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
	return "list"
}

func (c *command) Description() string {
	return "list of bus bookings"
}

func (c *command) Process(ctx context.Context, args string) string {
	params, err := commandPkg.CheckArguments(args, 2)
	if err != nil {
		return err.Error()
	}

	offset, err := commandPkg.StringToUint(params[0])
	if err != nil {
		return err.Error()
	}
	limit, err := commandPkg.StringToUint(params[1])
	if err != nil {
		return err.Error()
	}

	bbs, err := c.bb.List(ctx, offset, limit)
	if err != nil {
		return err.Error()
	}
	result := fmt.Sprint("id: route / date / seat")

	for _, bb := range bbs {
		result += fmt.Sprintf("\n%d: %s / %s / %d", bb.Id, bb.Route, bb.Date, bb.Seat)
	}
	return result
}
