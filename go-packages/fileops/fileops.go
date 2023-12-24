package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// write to a file
func writeFloatToFile(value float64, fileName string) {
	//convert the balance into a string of bytes
	valueText := fmt.Sprint(value)

	os.WriteFile(fileName, []byte(valueText), 0644)
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	//err will nil if we don't have an error
	if err != nil {
		return 1000, errors.New("Failed to find  file")
	}

	//can only use string - float, int will not work with a byte array
	valueText := string(data)
	//convert string to float
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}

	//we will return the balace and nil (meaning no error)
	return value, nil
}
