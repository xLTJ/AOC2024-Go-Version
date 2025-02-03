package lib

import "utils"

func Day9Part1(fileName string) (int, error) {
	diskMap, err := utils.ParseInput(fileName, parseLine)
	if err != nil {
		return 0, err
	}

	digitsPositions, emptyPositions := analyzeDiskMap(diskMap[0])

	compactedState := generateCompactedState(digitsPositions, emptyPositions)

	return calculateChecksum(compactedState), err
}

func parseLine(line string) ([]int, error) {
	return utils.StringToIntSlice(line, "")
}

// analyzeDiskMap returns two things based on the disk map:
//   - A map over positions in the initial state holding file blocks, with the value being the file block id
//   - A slice of positions that are empty
func analyzeDiskMap(diskMap []int) (map[int]int, []int) {
	var digitsPositions = map[int]int{}
	var emptyPositions []int

	counter := 0
	for i, element := range diskMap {
		if i%2 == 0 {
			for j := 0; j < element; j++ {
				digitsPositions[counter] = i / 2
				counter++
			}
			continue
		}

		for j := 0; j < element; j++ {
			emptyPositions = append(emptyPositions, counter)
			counter++
		}
	}

	return digitsPositions, emptyPositions
}

// generateCompactedState generates the compacted state based on the rules of the puzzle.
func generateCompactedState(digitsPositions map[int]int, emptyPositions []int) []int {
	rightMostPosition := len(digitsPositions) + len(emptyPositions) - 1

	var compactedState []int

	currentPosition := 0
	for _, emptyPosition := range emptyPositions {
		if rightMostPosition <= currentPosition {
			compactedState = append(compactedState, digitsPositions[currentPosition])
			break
		}

		for currentPosition < emptyPosition {
			compactedState = append(compactedState, digitsPositions[currentPosition])
			currentPosition++
		}

		_, ok := digitsPositions[rightMostPosition]
		for !ok {
			rightMostPosition--
			_, ok = digitsPositions[rightMostPosition]
		}

		if rightMostPosition <= currentPosition {
			break
		}

		compactedState = append(compactedState, digitsPositions[rightMostPosition])
		rightMostPosition--
		currentPosition++
	}

	return compactedState
}

// calculateChecksum calculates the checksum of a compacted state
func calculateChecksum(compactedState []int) int {
	checksum := 0
	for position, id := range compactedState {
		checksum += position * id
	}

	return checksum
}
