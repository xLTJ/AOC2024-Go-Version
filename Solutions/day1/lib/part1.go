package lib

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"utils"
)

type slicePair struct {
	left  []int
	right []int
}

func CalculateDistance(fileName string) (int, error) {
	// Parsing input
	var numberList slicePair
	var err error
	numberList.left, numberList.right, err = utils.ReadSplitAndParseLines(fileName, "   ", parseLine, parseLine)
	if err != nil {
		fmt.Println("Error reading and parsing file:", err)
		return 0, err
	}

	// Process the numbers
	sort.Ints(numberList.left)
	sort.Ints(numberList.right)

	distanceSum := 0
	for i, number := range numberList.left {
		distanceSum += int(math.Abs(float64(number - numberList.right[i])))
	}

	return distanceSum, nil
}

func parseLine(line string) (int, error) {
	number, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}

	return number, nil
}
