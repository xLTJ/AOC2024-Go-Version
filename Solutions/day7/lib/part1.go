package lib

import (
	"strconv"
	"strings"
	"utils"
)

type Calibration struct {
	result  int
	numbers []int
}

func CalculateTotalCalibration(fileName string) (int, error) {
	calibrations, err := utils.ParseInput(fileName, parseLine)
	if err != nil {
		return 0, err
	}

	result := 0
	for _, calibration := range calibrations {
		if isCalibrationValid(calibration) {
			result += calibration.result
		}
	}

	return result, err
}

func isCalibrationValid(calibration Calibration) bool {
	operatorAmount := len(calibration.numbers) - 1

	// uses a bitwise operator which is the same as 2^len
	for i := 0; i < (1 << operatorAmount); i++ {
		// converts i into its binary form, and makes sure the length is fixed to amount of operators by adding 0's
		// the binary representation basically tells the problem whether to use a + or * operator, that way we try every combination
		binary := strconv.FormatInt(int64(i), 2)
		if len(binary) < operatorAmount {
			binary = strings.Repeat("0", operatorAmount-len(binary)) + binary
		}
		sum := calibration.numbers[0]

		for j := range calibration.numbers {
			// as we check for j + 1 calibration number, we dont need to last one
			if j == operatorAmount {
				break
			}

			if binary[j] == '0' {
				sum += calibration.numbers[j+1]
			} else {
				sum *= calibration.numbers[j+1]
			}
		}

		if sum == calibration.result {
			return true
		}
	}
	return false
}

func parseLine(line string) (Calibration, error) {
	splitLine := strings.Split(line, ": ")

	var calibration Calibration
	var err error

	calibration.result, err = strconv.Atoi(splitLine[0])
	if err != nil {
		return Calibration{}, err
	}

	calibration.numbers, err = utils.StringToIntSlice(splitLine[1], " ")
	if err != nil {
		return Calibration{}, err
	}

	return calibration, nil
}
