package lib

import (
	"errors"
	"utils"
)

const (
	crossTarget        = "MAS"
	reverseCrossTarget = "SAM"
)

func CountMasCrosses(fileName string) (int, error) {
	var inputGrid runeGrid
	var err error

	inputGrid, err = utils.ParseInput(fileName, parseLine)
	if err != nil {
		return 0, err
	}

	counter := 0
	for y, row := range inputGrid {
		for x, _ := range row {
			hasMatch, err := inputGrid.checkCellForStringCross(coordinate{x, y}, crossTarget, reverseCrossTarget)
			if err != nil {
				return 0, err
			}
			if hasMatch {
				counter++
			}
		}
	}
	return counter, err
}

// checkCellForStringCross checks if a cell is the center of a cross of the target string
func (r runeGrid) checkCellForStringCross(startCell coordinate, targetString string, reverseTargetString string) (bool, error) {
	if len(targetString)%2 == 0 {
		return false, errors.New("target string can only have an odd amount of characters")
	}

	if r.getValue(startCell) != rune(targetString[len(targetString)/2]) {
		return false, nil
	}

	lineOneStartCell := coordinate{x: startCell.x - len(targetString)/2, y: startCell.y - len(targetString)/2}
	if !(r.checkLineForString(lineOneStartCell, targetString, coordinate{x: 1, y: 1}) || r.checkLineForString(lineOneStartCell, reverseTargetString, coordinate{x: 1, y: 1})) {
		return false, nil
	}

	lineTwoStartCell := coordinate{x: startCell.x - len(targetString)/2, y: startCell.y + len(targetString)/2}
	if !(r.checkLineForString(lineTwoStartCell, targetString, coordinate{x: 1, y: -1}) || r.checkLineForString(lineTwoStartCell, reverseTargetString, coordinate{x: 1, y: -1})) {
		return false, nil
	}
	return true, nil
}

func (r runeGrid) checkLineForString(startCell coordinate, targetString string, directionVector coordinate) bool {
	if r.isCellOutsideLimits(startCell) {
		return false
	}

	if r.getValue(startCell) != rune(targetString[0]) {
		return false
	}

	currentCell := startCell
	for i, character := range targetString {
		if i == 0 {
			continue
		}
		currentCell.x += directionVector.x
		currentCell.y += directionVector.y

		if r.isCellOutsideLimits(currentCell) {
			return false
		}

		if r.getValue(currentCell) != character {
			return false
		}
	}
	return true
}
