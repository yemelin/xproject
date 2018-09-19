package utils

import (
	"strconv"
)

// StringSliceToFloat64 converts slice of strings to slice of float64s
func StringSliceToFloat64(ss []string) ([]float64, error) {
	result := make([]float64, len(ss))
	for i, item := range ss {
		value, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return nil, err
		}
		result[i] = value
	}

	return result, nil
}
