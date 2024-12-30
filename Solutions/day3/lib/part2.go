package lib

import (
	"regexp"
	"strconv"
	"utils"
)

const (
	instructionEnable  = -1
	instructionDisable = -2
)

func CalculateInstructions2(fileName string) (int, error) {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)|(do\(\))|(don't\(\))`)

	instructions, err := utils.MatchAndParseInput(fileName, pattern, parseMatch2)

	if err != nil {
		return 0, err
	}

	sum := 0
	instructionsEnabled := true
	for _, instruction := range instructions {
		if instruction.x == instructionEnable {
			instructionsEnabled = true
		}
		if instruction.x == instructionDisable {
			instructionsEnabled = false
		}

		if instructionsEnabled {
			sum += instruction.x * instruction.y
		}
	}

	return sum, nil
}

func parseMatch2(match []string) (mulInstruction, error) {
	if match[0] == "don't()" {
		return mulInstruction{x: instructionDisable, y: 0}, nil
	}
	if match[0] == "do()" {
		return mulInstruction{x: instructionEnable, y: 0}, nil
	}

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
