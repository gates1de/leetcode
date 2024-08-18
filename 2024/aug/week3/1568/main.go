package main
import (
	"fmt"
)

var dirs = [][]int{{0,1},{1,0},{0,-1},{-1,0}}

type ArticulationPointInfo struct {
	HasPoint bool
	Time int
}

func minDays(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ap := ArticulationPointInfo{HasPoint: false, Time: 0}
	landCells := int(0)
	islandCount := int(0)

	discoveryTime := make([][]int, m)
	lowestReachable := make([][]int, m)
	parentCell := make([][]int, m)
	for i, _ := range discoveryTime {
		discoveryTime[i] = make([]int, n)
		lowestReachable[i] = make([]int, n)
		parentCell[i] = make([]int, n)

		for j, _ := range discoveryTime[i] {
			discoveryTime[i][j] = -1
			lowestReachable[i][j] = -1
			parentCell[i][j] = -1
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				landCells++
				if discoveryTime[i][j] == -1 {
					helper(grid, i, j, &discoveryTime, &lowestReachable, &parentCell, &ap)
					islandCount++
				}
			}
		}
	}

	if islandCount == 0 || islandCount >= 2 {
		return 0
	}
	if landCells == 1 {
		return 1
	}
	if ap.HasPoint {
		return 1
	}
	return 2
}

func helper(
	grid [][]int,
	row int,
	col int,
	discoveryTime *[][]int,
	lowestReachable *[][]int,
	parentCell *[][]int,
	ap *ArticulationPointInfo,
) {
	n := len(grid[0])
	(*discoveryTime)[row][col] = ap.Time
	ap.Time++
	(*lowestReachable)[row][col] = (*discoveryTime)[row][col]
	children := int(0)

	for _, dir := range dirs {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if isValidLandCell(grid, newRow, newCol) {
			if (*discoveryTime)[newRow][newCol] == -1 {
				children++
				(*parentCell)[newRow][newCol] = row * n + col
				helper(grid, newRow, newCol, discoveryTime, lowestReachable, parentCell, ap)
				(*lowestReachable)[row][col] = min((*lowestReachable)[row][col], (*lowestReachable)[newRow][newCol])

				if (*lowestReachable)[newRow][newCol] >= (*discoveryTime)[row][col] && (*parentCell)[row][col] != -1 {
					ap.HasPoint = true
				}
			} else if newRow * n + newCol != (*parentCell)[row][col] {
				(*lowestReachable)[row][col] = min((*lowestReachable)[row][col], (*discoveryTime)[newRow][newCol])
			}
		}
	}

	if (*parentCell)[row][col] == -1 && children > 1 {
		ap.HasPoint = true
	}
}

func isValidLandCell(grid [][]int, row int, col int) bool {
	m := len(grid)
	n := len(grid[0])

	return row >= 0 && col >= 0 && row < m && col < n && grid[row][col] == 1
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// grid := [][]int{{0,1,1,0},{0,1,1,0},{0,0,0,0}}

	// result: 2
	grid := [][]int{{1,1}}

	// result: 
	// grid := [][]int{}

	result := minDays(grid)
	fmt.Printf("result = %v\n", result)
}

