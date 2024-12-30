package lib

import (
	"math"
	"strconv"
	"strings"
	"utils"
)

const maxDifference = 3

// ParseLine parses a line in the format of the input file to a slice of int's
func ParseLine(line string) ([]int, error) {
	var parsedLine []int

	for _, number := range strings.Split(line, " ") {
		parsedNumber, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		parsedLine = append(parsedLine, parsedNumber)
	}

	return parsedLine, nil
}

// CountSafeReports counts how many reports are safe from an input file
func CountSafeReports(fileName string) (int, error) {
	reports, err := utils.ReadAndParseLines(fileName, ParseLine)
	if err != nil {
		return 0, err
	}

	safeReports := 0
	for _, report := range reports {
		isSafe, err := CheckSafety(report)
		if err != nil {
			return 0, err
		}

		if isSafe {
			safeReports++
		}
	}
	return safeReports, nil
}

// CheckSafety checks if a report is safe or not based on the given requirements
func CheckSafety(report []int) (bool, error) {
	isIncreasing := report[1] > report[0]

	for i, number := range report {
		if i == 0 {
			continue
		}

		difference := number - report[i-1]
		if math.Abs(float64(difference)) < 1 || math.Abs(float64(difference)) > maxDifference {
			return false, nil
		}

		if (isIncreasing && difference <= 0) || (!isIncreasing && difference >= 0) {
			return false, nil
		}
	}
	return true, nil
}
