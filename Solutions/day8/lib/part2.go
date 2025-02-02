package lib

import "utils"

func (m MapGrid) getAllAntinodes2(cellsToCheck map[rune][]Coordinate) map[Coordinate]bool {
	var antinodeMap = map[Coordinate]bool{}
	for _, cells := range cellsToCheck {
		m.getAntinodesForFrequency2(cells, antinodeMap)
	}

	return antinodeMap
}

func (m MapGrid) getAntinodesForFrequency2(cells []Coordinate, antinodeMap map[Coordinate]bool) {
	for i, cellA := range cells {
		for _, cellB := range cells[i+1:] {
			m.getAntinodesForCells2(cellA, cellB, antinodeMap)
		}
	}
}

func (m MapGrid) getAntinodesForCells2(cellA, cellB Coordinate, antinodeMap map[Coordinate]bool) {
	antinodeMap[cellA] = true
	antinodeMap[cellB] = true

	coordinateDifference := Coordinate{x: cellA.x - cellB.x, y: cellA.y - cellB.y}

	tempAntiNode := Coordinate{x: cellA.x + coordinateDifference.x, y: cellA.y + coordinateDifference.y}

	for !m.isCellOutsideGrid(tempAntiNode) {
		antinodeMap[tempAntiNode] = true
		tempAntiNode.x += coordinateDifference.x
		tempAntiNode.y += coordinateDifference.y
	}

	tempAntiNode = Coordinate{x: cellB.x - coordinateDifference.x, y: cellB.y - coordinateDifference.y}

	for !m.isCellOutsideGrid(tempAntiNode) {
		antinodeMap[tempAntiNode] = true
		tempAntiNode.x -= coordinateDifference.x
		tempAntiNode.y -= coordinateDifference.y
	}
}

func CountAntinodes2(fileName string) (int, error) {
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

	antinodeMap := antennaMap.getAllAntinodes2(cellsToCheck)
	return len(antinodeMap), nil
}
