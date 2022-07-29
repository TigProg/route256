package tools

import (
	"strconv"

	"github.com/pkg/errors"
)

var NonNumericString = errors.New("bus booking does not exist")

func StringToUint(s string) (uint, error) {
	resultUint64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, errors.Wrapf(NonNumericString, s)
	}
	return uint(resultUint64), nil
}
