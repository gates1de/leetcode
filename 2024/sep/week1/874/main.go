package main

import (
	"fmt"
)

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func robotSim(commands []int, obstacles [][]int) int {
	result := int(0)
	obstaclesMap := make(map[[2]int]bool)
	for _, obstacle := range obstacles {
		obstaclesMap[[2]int{obstacle[0], obstacle[1]}] = true
	}

	currentPosition := [2]int{0, 0}
	currentDirection := int(0) // 0: North, 1: East, 2: South, 3: West

	for _, command := range commands {
		if command == -1 {
			currentDirection = (currentDirection + 1) % 4
			continue
		}
		if command == -2 {
			currentDirection = (currentDirection + 3) % 4
			continue
		}

		direction := dirs[currentDirection]
		for step := 0; step < command; step++ {
			nextX := currentPosition[0] + direction[0]
			nextY := currentPosition[1] + direction[1]
			if obstaclesMap[[2]int{nextX, nextY}] {
				break
			}
			currentPosition[0] = nextX
			currentPosition[1] = nextY
		}

		result = max(
			result,
			currentPosition[0] * currentPosition[0] + currentPosition[1] * currentPosition[1],
		)
	}

	return result
}

func hashCoordinates(x, y int) int {
	return x + 1000000 * y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// result: 25
	// commands := []int{4, -1, 3}
	// obstacles := [][]int{}

	// result: 65
	// commands := []int{4, -1, 4, -2, 4}
	// obstacles := [][]int{{2, 4}}

	// result: 36
	commands := []int{6, -1, -1, 6}
	obstacles := [][]int{}

	// result:
	// commands := []int{}
	// obstacles := [][]int{}

	result := robotSim(commands, obstacles)
	fmt.Printf("result = %v\n", result)
}
