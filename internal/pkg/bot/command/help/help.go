package help

import (
	"context"
	"fmt"

	commandPkg "gitlab.ozon.dev/tigprog/bus_booking/internal/pkg/bot/command"
)

func New(extendedMap map[string]string) commandPkg.Interface {
	if extendedMap == nil {
		extendedMap = map[string]string{}
	}
	return &command{
		extended: extendedMap,
	}
}

type command struct {
	extended map[string]string
}

func (c *command) Name() string {
	return "help"
}

func (c *command) Description() string {
	return "list of commands"
}

func (c *command) Process(_ context.Context, _ string) string {
	result := fmt.Sprintf("/%s - %s", c.Name(), c.Description())
	for cmd, description := range c.extended {
		result += fmt.Sprintf("\n/%s - %s", cmd, description)
	}
	return result
}
