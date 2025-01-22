package lib

import (
	"errors"
	"utils"
)

const (
	guardRune     = '^'
	wallRune      = '#'
	emptyRune     = '.'
	maxIterations = 100000
)

type MapGrid [][]rune

type Coordinate struct {
	x int
	y int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

var Directions = map[Direction]Coordinate{
	Up:    {0, -1},
	Right: {1, 0},
	Down:  {0, 1},
	Left:  {-1, 0},
}

type Guard struct {
	coordinate Coordinate
	direction  Direction
}

func (g *Guard) turnRight() {
	g.direction = (g.direction + 1) % 4
}

func (g *Guard) getNextCell() Coordinate {
	return Coordinate{
		x: g.coordinate.x + Directions[g.direction].x,
		y: g.coordinate.y + Directions[g.direction].y,
	}
}

func (g *Guard) move() {
	g.coordinate.x += Directions[g.direction].x
	g.coordinate.y += Directions[g.direction].y
}

func (m MapGrid) getCellValue(cell Coordinate) rune {
	return m[cell.y][cell.x]
}

func (m MapGrid) isCellOutsideGrid(cell Coordinate) bool {
	return (cell.x < 0 || cell.x > len(m[0])-1) || (cell.y < 0 || cell.y > len(m)-1)
}

func (m MapGrid) getGuardLocation() (Coordinate, error) {
	for y, row := range m {
		for x, cell := range row {
			if cell == guardRune {
				return Coordinate{x, y}, nil
			}
		}
	}
	return Coordinate{}, errors.New("Guard not found")
}

func CountCellsPassed(fileName string) (int, error) {
	var mapGrid MapGrid
	var err error
	mapGrid, err = utils.ParseInput(fileName, parseLine)
	if err != nil {
		return 0, err
	}

	guardLocation, err := mapGrid.getGuardLocation()
	if err != nil {
		return 0, err
	}

	guard := Guard{direction: Up, coordinate: guardLocation}
	passedCells := simulateGuardPath(mapGrid, guard)
	return len(passedCells), nil
}

func simulateGuardPath(mapGrid MapGrid, guard Guard) map[Coordinate]bool {
	var passedCells = map[Coordinate]bool{}

	for i := 0; i < maxIterations; i++ {
		nextCell := guard.getNextCell()

		if mapGrid.isCellOutsideGrid(nextCell) {
			break
		}

		if mapGrid.getCellValue(nextCell) == wallRune {
			guard.turnRight()
			continue
		}

		guard.move()
		passedCells[guard.coordinate] = true
	}

	return passedCells
}

func parseLine(line string) ([]rune, error) {
	return []rune(line), nil
}
