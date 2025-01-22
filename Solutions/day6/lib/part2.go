package lib

import "utils"

type GuardState struct {
	coordinate Coordinate
	direction  Direction
}

func CountLoopCreatingObstacles(fileName string) (int, error) {
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

	originalGuard := Guard{direction: Up, coordinate: guardLocation}

	originalPassedCells := simulateGuardPath2(mapGrid, originalGuard)

	loopAmount := 0

	for y, row := range mapGrid {
		for x, cell := range row {
			if cell == wallRune || cell == guardRune {
				continue
			}

			val, ok := originalPassedCells[Coordinate{x, y}]
			if !ok {
				continue
			}

			newGuard := Guard{direction: val.direction, coordinate: val.coordinate}

			mapGrid[y][x] = wallRune
			if doesGuardLoop(mapGrid, newGuard) {
				loopAmount++
			}
			mapGrid[y][x] = emptyRune
		}
	}

	return loopAmount, nil
}

func doesGuardLoop(mapGrid MapGrid, guard Guard) bool {
	var passedGuardStates = map[GuardState]bool{}

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
		if passedGuardStates[GuardState(guard)] {
			return true
		}
		passedGuardStates[GuardState(guard)] = true
	}
	return false
}

// simulateGuardPath2 does the same as simulateGuardPath, but for each coordinate in the map its value is the state
// of the previous cell. That way we can start the guard at that location and not have to run the entire simulation before
func simulateGuardPath2(mapGrid MapGrid, guard Guard) map[Coordinate]GuardState {
	var passedCells = map[Coordinate]GuardState{}

	for i := 0; i < maxIterations; i++ {
		nextCell := guard.getNextCell()

		if mapGrid.isCellOutsideGrid(nextCell) {
			break
		}

		if mapGrid.getCellValue(nextCell) == wallRune {
			guard.turnRight()
			continue
		}

		_, ok := passedCells[nextCell]

		if !ok {
			passedCells[nextCell] = GuardState(guard)
		}
		guard.move()
	}

	return passedCells
}
