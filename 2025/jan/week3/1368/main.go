package main
import (
	"fmt"
	"math"
)

func minCost(grid [][]int) int {
	dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}

	m := len(grid)
	n := len(grid[0])

	minCost := make([][]int, m)
	for i, _ := range minCost {
		minCost[i] = make([]int, n)

		for j, _ := range minCost[i] {
			minCost[i][j] = math.MaxInt32
		}
	}

	deque := make([][]int, 0)
	deque = append(deque, []int{0, 0})
	minCost[0][0] = 0

	for len(deque) > 0 {
		current := deque[0]
		deque = deque[1:]
		row := current[0]
		col := current[1]

		for i, dir := range dirs {
			newRow := row + dir[0]
			newCol := col + dir[1]
			cost := int(0)
			if grid[row][col] != i + 1 {
				cost = 1
			}

			if isValid(newRow, newCol, m, n) && minCost[row][col] + cost < minCost[newRow][newCol] {
				minCost[newRow][newCol] = minCost[row][col] + cost

				if cost == 1 {
					deque = append(deque, []int{newRow, newCol})
				} else {
					deque = append(deque, []int{0, 0})
					copy(deque[1:], deque)
					deque[0] = []int{newRow, newCol}
				}
			}
		}
	}

	return minCost[m - 1][n - 1]
}

func isValid(row int, col int, numOfRows int, numOfCols int) bool {
	return row >= 0 && row < numOfRows && col >= 0 && col < numOfCols
}

func main() {
	// result: 3
	// grid := [][]int{{1,1,1,1},{2,2,2,2},{1,1,1,1},{2,2,2,2}}

	// result: 0
	// grid := [][]int{{1,1,3},{3,2,2},{1,1,4}}

	// result: 1
	grid := [][]int{{1,2},{4,3}}

	// result: 0
	// grid := [][]int{}

	result := minCost(grid)
	fmt.Printf("result = %v\n", result)
}

