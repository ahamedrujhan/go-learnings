package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(lines []string) ([]float64, error) {

	result := make([]float64, len(lines))

	for i, line := range lines {
		floatVal, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New(fmt.Sprint("Error converting ", line, " to float."))
		}
		result[i] = floatVal
	}

	return result, nil
}
