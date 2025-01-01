package lib

import "utils"

const (
	target = "XMAS"
)

type coordinate struct {
	x int
	y int
}

type runeGrid [][]rune

var directionVectors = []coordinate{
	{0, 1},
	{1, 0},
	{1, 1},
	{-1, 1},
	{0, -1},
	{-1, 0},
	{-1, -1},
	{1, -1},
}

func CountXmas(fileName string) (int, error) {
	var inputGrid runeGrid
	var err error

	inputGrid, err = utils.ParseInput(fileName, parseLine)
	if err != nil {
		return 0, err
	}
	return inputGrid.checkGridForString(target, directionVectors), nil
}

func (r runeGrid) getValue(cell coordinate) rune {
	return r[cell.y][cell.x]
}

func (r runeGrid) isCellOutsideLimits(cell coordinate) bool {
	return (cell.x < 0 || cell.x > len(r[0])-1) || (cell.y < 0 || cell.y > len(r)-1)
}

func (r runeGrid) checkGridForString(targetString string, validDirections []coordinate) int {
	counter := 0
	for y, row := range r {
		for x, _ := range row {
			counter += r.checkCellForString(coordinate{x: x, y: y}, targetString, validDirections)
		}
	}
	return counter
}

func (r runeGrid) checkCellForString(startCell coordinate, targetString string, validDirections []coordinate) int {
	if r.getValue(startCell) != rune(targetString[0]) {
		return 0
	}

	matches := 0
	for _, directionVector := range validDirections {
		currentCell := startCell
		match := true

		for i, character := range targetString {
			// skip the first character cus we already did that one
			if i == 0 {
				continue
			}
			currentCell.x += directionVector.x
			currentCell.y += directionVector.y

			if r.isCellOutsideLimits(currentCell) {
				match = false
				break
			}

			if r.getValue(currentCell) != character {
				match = false
				break
			}
		}

		if match {
			matches++
		}
	}

	return matches
}

func parseLine(line string) ([]rune, error) {
	return []rune(line), nil
}
