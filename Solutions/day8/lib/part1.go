package lib

import "utils"

const emptyCell = '.'

type MapGrid [][]rune

type Coordinate struct {
	x int
	y int
}

func (m MapGrid) getCellValue(cell Coordinate) rune {
	return m[cell.y][cell.x]
}

func (m MapGrid) isCellOutsideGrid(cell Coordinate) bool {
	return (cell.x < 0 || cell.x > len(m[0])-1) || (cell.y < 0 || cell.y > len(m)-1)
}

// getAllAntinodes takes a map of all the cells to check (grouped by frequency) and returns a map of all the coordinates
// with antinodes on them.
func (m MapGrid) getAllAntinodes(cellsToCheck map[rune][]Coordinate) map[Coordinate]bool {
	var antinodeMap = map[Coordinate]bool{}
	for _, cells := range cellsToCheck {
		m.getAntinodesForFrequency(cells, antinodeMap)
	}

	return antinodeMap
}

// getAntinodesForFrequency goes through every cell for a frequency (given as a slice of Coordinates) and adds all their
// antinodes to the antinode map.
func (m MapGrid) getAntinodesForFrequency(cells []Coordinate, antinodeMap map[Coordinate]bool) {
	for i, cellA := range cells {
		for _, cellB := range cells[i+1:] {
			antinodeA, antinodeB := m.getAntinodesForCells(cellA, cellB)
			if !m.isCellOutsideGrid(antinodeA) {
				antinodeMap[antinodeA] = true
			}
			if !m.isCellOutsideGrid(antinodeB) {
				antinodeMap[antinodeB] = true
			}
		}
	}
}

// getAntinodesForCells returns the antinodes for two cells
func (m MapGrid) getAntinodesForCells(cellA, cellB Coordinate) (Coordinate, Coordinate) {
	coordinateDifference := Coordinate{x: cellA.x - cellB.x, y: cellA.y - cellB.y}
	antinodeA := Coordinate{x: cellA.x + coordinateDifference.x, y: cellA.y + coordinateDifference.y}
	antinodeB := Coordinate{x: cellB.x - coordinateDifference.x, y: cellB.y - coordinateDifference.y}

	return antinodeA, antinodeB
}

// CountAntinodes is the main function that counds how many antinodes there are
func CountAntinodes(fileName string) (int, error) {
	var antennaMap MapGrid
	antennaMap, err := utils.ParseInput(fileName, parseLine)
	if err != nil {
		return 0, err
	}

	var cellsToCheck = map[rune][]Coordinate{}

	// gets all the non-empty cells and groups them by frequency in a map
	for y, row := range antennaMap {
		for x, cell := range row {
			if cell == emptyCell {
				continue
			}

			_, ok := cellsToCheck[cell]
			if !ok {
				cellsToCheck[cell] = []Coordinate{{x, y}}
			} else {
				cellsToCheck[cell] = append(cellsToCheck[cell], Coordinate{x, y})
			}
		}
	}

	antinodeMap := antennaMap.getAllAntinodes(cellsToCheck)
	return len(antinodeMap), nil
}

func parseLine(line string) ([]rune, error) {
	return []rune(line), nil
}
