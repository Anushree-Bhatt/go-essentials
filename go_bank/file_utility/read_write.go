package file_utility

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fname string) (float64, error) {
	data_byte, err := os.ReadFile(fname)

	if err != nil {
		custom_err := errors.New("file not found, hence defaulting value to 0")
		return 0, custom_err
	}

	data_txt := string(data_byte)
	data, err := strconv.ParseFloat(data_txt, 64)

	if err != nil {
		custom_err := errors.New("file doesn't contain proper data, hence defaulting value to 0")
		return 0, custom_err
	}

	return data, nil
}

func WriteFloatToFile(fname string, data float64) {
	data_txt := fmt.Sprint(data)
	err := os.WriteFile(fname, []byte(data_txt), 0644)

	if err != nil {
		fmt.Print("Writing file error:", err)
	}
}
