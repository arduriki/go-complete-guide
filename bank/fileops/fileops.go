package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// getBalanceFromFile reads text from a file.
func GetFloatFromFile(filename string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}
	// Convert from file to text.
	valueText := string(data)
	// Convert from str to float64.
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}

	return value, nil
}

// writeFloatToFile writes text into a file.
func WriteFloatToFile(value float64, fileName string) {
	// To create a collection of bytes.
	valueText := fmt.Sprint(value)
	// Indicate in which file you want to store the byte array and set the permissions.
	os.WriteFile(fileName, []byte(valueText), 0644)
}