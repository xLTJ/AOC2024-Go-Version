package lib

import (
	"math"
	"strconv"
	"strings"
	"utils"
)

func CalculateTotalCalibration2(fileName string) (int, error) {
	calibrations, err := utils.ParseInput(fileName, parseLine)
	if err != nil {
		return 0, err
	}

	result := 0
	for _, calibration := range calibrations {
		if isCalibrationValid2(calibration) {
			result += calibration.result
		}
	}

	return result, err
}

func isCalibrationValid2(calibration Calibration) bool {
	operatorAmount := len(calibration.numbers) - 1

	for i := 0; i < int(math.Pow(3, float64(operatorAmount))); i++ {

		// same as in part 1, but with a base 3 instead cus we have 3 operators now
		trinary := strconv.FormatInt(int64(i), 3)
		if len(trinary) < operatorAmount {
			trinary = strings.Repeat("0", operatorAmount-len(trinary)) + trinary
		}
		sum := calibration.numbers[0]

		for j := range calibration.numbers {
			// as we check for j + 1 calibration number, we dont need to last one
			if j == operatorAmount {
				break
			}

			switch trinary[j] {
			case '0':
				sum += calibration.numbers[j+1]
			case '1':
				sum *= calibration.numbers[j+1]
			case '2':
				sum = concatenateInts(sum, calibration.numbers[j+1])
			}
		}

		if sum == calibration.result {
			return true
		}
	}
	return false
}

func concatenateInts(a int, b int) int {
	multiplier := 1
	tempB := b

	for tempB > 0 {
		multiplier *= 10
		tempB /= 10
	}
	if b == 0 {
		multiplier = 10
	}

	return a*multiplier + b
}
