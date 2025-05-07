package number

import (
	"errors"
	"math"
)

func SafeIntToInt32(value int) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, errors.New("value out of int32 range")
	}
	return int32(value), nil
}
