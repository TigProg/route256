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

func (c *command) Process(ctx context.Context, _ string) string {
	bbs := c.bb.List(ctx)
	result := fmt.Sprint("id: route / date / seat")

	for _, bb := range bbs {
		result += fmt.Sprintf("\n%d: %s / %s / %d", bb.Id, bb.Route, bb.Date, bb.Seat)
	}
	return result
}
