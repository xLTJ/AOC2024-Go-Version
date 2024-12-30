package lib

import (
	"fmt"
	"utils"
)

func CalculateSimilarity(fileName string) (int, error) {
	var numberList slicePair
	var err error
	numberList.left, numberList.right, err = utils.SplitAndParseInput(fileName, "   ", parseLine, parseLine)
	if err != nil {
		fmt.Println("Error reading and parsing file:", err)
		return 0, err
	}

	var rightListFrequencyMap = map[int]int{}
	for _, number := range numberList.right {
		rightListFrequencyMap[number]++
	}

	similarityScore := 0
	for _, number := range numberList.left {
		similarityScore += number * rightListFrequencyMap[number]
	}

	return similarityScore, nil
}
