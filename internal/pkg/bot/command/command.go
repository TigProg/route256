package command

import (
	"context"
	"fmt"
	"strings"
)

type Interface interface {
	Name() string
	Description() string
	Process(ctx context.Context, args string) string
}

func CheckArguments(args string, expected int) ([]string, error) {
	var params []string
	if args == "" {
		params = make([]string, 0)
	} else {
		params = strings.Split(args, " ")
	}

	if len(params) != expected {
		return nil, fmt.Errorf(
			"incorrect number of arguments, expected %d, got %d",
			expected, len(params),
		)
	}
	return params, nil
}
