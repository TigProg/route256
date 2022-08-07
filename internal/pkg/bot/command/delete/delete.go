package delete

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
	return "delete"
}

func (c *command) Description() string {
	return "delete bus booking by id"
}

func (c *command) Process(ctx context.Context, args string) string {
	params, err := commandPkg.CheckArguments(args, 1)
	if err != nil {
		return err.Error()
	}

	id, err := toolsPkg.StringToUint(params[0])
	if err != nil {
		return err.Error()
	}

	if err := c.bb.Delete(ctx, id); err != nil {
		return err.Error()
	}
	return fmt.Sprintf("[SUCCESS]\nBus booking seat deleted successfully\nbooking number: %d", id)
}
