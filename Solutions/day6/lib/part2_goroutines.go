package lib

import (
	"runtime"
	"sync"
	"utils"
)

func CountLoopCreatingObstaclesGoroutine(fileName string) (int, error) {
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

	// taskQueue is a queue of obstacle coordinates to check. Workers read from this channel
	taskQueue := make(chan Coordinate)

	// results is a channel for the workers to send the result back bout whether there's a loop
	results := make(chan int)
	workersAmount := runtime.NumCPU()
	var wg sync.WaitGroup

	for i := 0; i < workersAmount; i++ {
		wg.Add(1)
		go worker(mapGrid, originalPassedCells, taskQueue, results, &wg)
	}

	// populates the task queue, basically the same loop as in the original version, but only the part where we get the
	// valid coordinates, which are then added to the queue channel for the workers to simulate.
	go func() {
		defer close(taskQueue)
		for y, row := range mapGrid {
			for x, cell := range row {
				if cell == wallRune || cell == guardRune {
					continue
				}

				_, ok := originalPassedCells[Coordinate{x, y}]
				if !ok {
					continue
				}

				taskQueue <- Coordinate{x, y}
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	loopAmount := 0
	for result := range results {
		loopAmount += result
	}

	return loopAmount, nil
}

func worker(mapGrid MapGrid, originalPassedStates map[Coordinate]GuardState, taskQueue <-chan Coordinate, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for obstacle := range taskQueue {
		val := originalPassedStates[obstacle]
		newGuard := Guard{direction: val.direction, coordinate: val.coordinate}

		if doesGuardLoop2(mapGrid, newGuard, obstacle) {
			result <- 1
		} else {
			result <- 0
		}
	}
}

// doesGuardLoop2 is the same as the original version but it gets the obstacle as a separate coordinate instead
// This way the workers are able to do stuff at the same time, as otherwise the new map would be modified multiple times
// or we would have to create a deep copy of the map every time (very inefficient)
func doesGuardLoop2(mapGrid MapGrid, guard Guard, obstacle Coordinate) bool {
	var passedGuardStates = map[GuardState]bool{}

	for i := 0; i < maxIterations; i++ {
		nextCell := guard.getNextCell()

		if mapGrid.isCellOutsideGrid(nextCell) {
			break
		}

		if mapGrid.getCellValue(nextCell) == wallRune || nextCell == obstacle {
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
