package main
import (
	"fmt"
)

func highestPeak(isWater [][]int) [][]int {
	m := len(isWater)
	n := len(isWater[0])

	result := make([][]int, m)
	for i, _ := range result {
		result[i] = make([]int, n)
		for j, _ := range result[i] {
			result[i][j] = m * n
		}
	}


	for i, _ := range result {
		for j, _ := range result[i] {
			if isWater[i][j] == 1 {
				result[i][j] = 0
			}
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			minNeighborDistance := m * n

			neighborRow := row - 1
			neighborCol := col
			if isValid(neighborRow, neighborCol, m, n) {
				minNeighborDistance = min(minNeighborDistance, result[neighborRow][neighborCol])
			}

			neighborRow = row
			neighborCol = col - 1
			if isValid(neighborRow, neighborCol, m, n) {
				minNeighborDistance = min(minNeighborDistance, result[neighborRow][neighborCol])
			}

			result[row][col] = min(result[row][col], minNeighborDistance + 1)
		}
	}

	for row := m - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			minNeighborDistance := m * n

			neighborRow := row + 1
			neighborCol := col
			if isValid(neighborRow, neighborCol, m, n) {
				minNeighborDistance = min(minNeighborDistance, result[neighborRow][neighborCol])
			}

			neighborRow = row
			neighborCol = col + 1
			if isValid(neighborRow, neighborCol, m, n) {
				minNeighborDistance = min(minNeighborDistance, result[neighborRow][neighborCol])
			}

			result[row][col] = min(result[row][col], minNeighborDistance + 1)
		}
	}

	return result
}

func isValid(row int, col int, numOfRows int, numOfCols int) bool {
	return row >= 0 && row < numOfRows && col >= 0 && col < numOfCols
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [[1,0],[2,1]]
	// isWater := [][]int{{0,1},{0,0}}

	// result: [[1,1,0],[0,1,1],[1,2,2]]
	isWater := [][]int{{0,0,1},{1,0,0},{0,0,0}}

	// result: []
	// isWater := [][]int{}

	result := highestPeak(isWater)
	fmt.Printf("result = %v\n", result)
}

