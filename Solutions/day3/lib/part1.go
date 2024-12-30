package lib

import (
	"regexp"
	"strconv"
	"utils"
)

type mulInstruction struct {
	x int
	y int
}

func CalculateInstructions(fileName string) (int, error) {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	instructions, err := utils.MatchAndParseInput(fileName, pattern, parseMatch)

	if err != nil {
		return 0, nil
	}

	sum := 0
	for _, instruction := range instructions {
		sum += instruction.x * instruction.y
	}

	return sum, nil
}

func parseMatch(match []string) (mulInstruction, error) {
	x, err := strconv.Atoi(match[1])
	if err != nil {
		return mulInstruction{}, err
	}

	y, err := strconv.Atoi(match[2])
	if err != nil {
		return mulInstruction{}, err
	}

	return mulInstruction{x: x, y: y}, nil
}
