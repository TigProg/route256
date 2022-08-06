package tools

import (
	"strconv"

	"github.com/pkg/errors"
)

var ErrNonNumericString = errors.New("string must be a non-negative number")

func StringToUint(s string) (uint, error) {
	resultUint64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, errors.Wrapf(ErrNonNumericString, s)
	}
	return uint(resultUint64), nil
}
