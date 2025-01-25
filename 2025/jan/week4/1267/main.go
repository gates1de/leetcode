package main
import (
	"fmt"
)

func countServers(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	result := int(0)
	rowCounts := make([]int, n)
	lastServerInCol := make([]int, m)
	for i, _ := range lastServerInCol {
		lastServerInCol[i] = -1
	}

	for row := 0; row < m; row++ {
		serverCountInRow := int(0)

		for col := 0; col < n; col++ {
			if grid[row][col] == 1 {
				serverCountInRow++
				rowCounts[col]++
				lastServerInCol[row] = col
			}
		}

		if serverCountInRow != 1 {
			result += serverCountInRow
			lastServerInCol[row] = -1
		}
	}

	for row := 0; row < m; row++ {
		if lastServerInCol[row] != -1 && rowCounts[lastServerInCol[row]] > 1 {
			result++
		}
	}

	return result
}

func main() {
	// result: 0
	// grid := [][]int{{1,0},{0,1}}

	// result: 3
	// grid := [][]int{{1,0},{1,1}}

	// result: 4
	grid := [][]int{{1,1,0,0},{0,0,1,0},{0,0,1,0},{0,0,0,1}}

	// result: 
	// grid := [][]int{}

	result := countServers(grid)
	fmt.Printf("result = %v\n", result)
}

