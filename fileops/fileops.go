package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(pathToFile string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(pathToFile, []byte(valueText), 0644)
}

func GetFloatFromFile(pathToFile string) (float64, error) {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Cannot convert it ")
	}

	return balance, nil
}
