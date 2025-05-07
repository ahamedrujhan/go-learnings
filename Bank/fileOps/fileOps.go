package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetValueFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	// error handling if the balance.txt not found
	if err != nil {
		return 1000, errors.New("Failed to find a file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	// error handling if the float conversion found
	if err != nil {
		return 1000, errors.New("Failed to parse value from file")
	}
	return value, nil
}

func WriteValueFromFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)

	err := os.WriteFile(fileName, []byte(valueText), 0644)

	if err != nil {
		panic("Failed to write value to file")
	}
}
