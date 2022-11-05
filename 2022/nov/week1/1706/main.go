package main
import (
	"fmt"
)

func findBall(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	ballPositions := make([][]int, m + 1)

	for i, _ := range ballPositions {
		ballPositions[i] = make([]int, n)
		for j, _ := range ballPositions[i] {
			if i == 0 {
				ballPositions[i][j] = j
			} else {
				ballPositions[i][j] = 10000
			}
		}
	}

	for i, _ := range grid {
		for j, _ := range grid[i] {
			span := grid[i][j]
			ballPosition := ballPositions[i][j]
			if ballPosition == 10000 {
				continue
			}

			if span == 1 {
				if j + 1 >= n || grid[i][j + 1] == -1 {
					continue
				}

				ballPositions[i + 1][j + 1] = ballPosition
			} else {
				if j - 1 < 0 || grid[i][j - 1] == 1 {
					continue
				}

				ballPositions[i + 1][j - 1] = ballPosition
			}
		}
	}

	result := make([]int, n)
	for i, _ := range result {
		result[i] = -1
	}

	for i, position := range ballPositions[m] {
		if position == 10000 {
			continue
		}
		result[position] = i
	}
	return result
}

func main() {
	// result: [1,-1,-1,-1,-1]
	// grid := [][]int{
	// 	{1,1,1,-1,-1},
	// 	{1,1,1,-1,-1},
	// 	{-1,-1,-1,1,1},
	// 	{1,1,1,1,-1},
	// 	{-1,-1,-1,-1,-1},
	// }

	// result: [-1]
	// grid := [][]int{{-1}}

	// result: [0,1,2,3,4,-1]
	grid := [][]int{
		{1,1,1,1,1,1},
		{-1,-1,-1,-1,-1,-1},
		{1,1,1,1,1,1},
		{-1,-1,-1,-1,-1,-1},
	}

	// result: 
	// grid := [][]int{}

	result := findBall(grid)
	fmt.Printf("result = %v\n", result)
}

