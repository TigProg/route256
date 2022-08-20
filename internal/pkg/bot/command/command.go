package command

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrNonNumericString        = errors.New("string must be a non-negative number")
	ErrIncorrectArgumentNumber = errors.New("")
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
		return nil, errors.Wrapf(
			ErrIncorrectArgumentNumber,
			"incorrect number of arguments, expected %d, got %d",
			expected, len(params),
		)
	}
	return params, nil
}

func StringToUint(s string) (uint, error) {
	resultUint64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, errors.Wrap(ErrNonNumericString, s)
	}
	return uint(resultUint64), nil
}
