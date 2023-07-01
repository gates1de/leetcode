package main
import (
	"fmt"
)

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func latestDayToCross(row int, col int, cells [][]int) int {
	left := int(1)
	right := row * col

	for left < right {
		mid := right - (right - left) / 2
		if canCross(row, col, mid, cells) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func dfs(r int, c int, row int, col int, grid [][]int) bool {
	if r < 0 || r >= row || c < 0 || c >= col || grid[r][c] != 0 {
		return false
	}

	if r == row - 1 {
		return true
	}

	grid[r][c] = -1
	for _, dir := range dirs {
		newR := r + dir[0]
		newC := c + dir[1]

		if dfs(newR, newC, row, col, grid) {
			return true
		}
	}

	return false
}

func canCross(row int, col int, day int, cells [][]int) bool {
	grid := make([][]int, row)
	for i, _ := range grid {
		grid[i] = make([]int, col)
	}

	for i := 0; i < day; i++ {
		r := cells[i][0] - 1
		c := cells[i][1] - 1
		grid[r][c] = 1
	}

	for i := 0; i < day; i++ {
		grid[cells[i][0] - 1][cells[i][1] - 1] = 1
	}

	for i := 0; i < col; i++ {
		if grid[0][i] == 0 && dfs(0, i, row, col, grid) {
			return true
		}
	}

	return false
}

func main() {
	// result: 2
	// row := int(2)
	// col := int(2)
	// cells := [][]int{{1,1},{2,1},{1,2},{2,2}}

	// result: 1
	// row := int(2)
	// col := int(2)
	// cells := [][]int{{1,1},{1,2},{2,1},{2,2}}

	// result: 3
	row := int(3)
	col := int(3)
	cells := [][]int{{1,2},{2,1},{3,3},{2,2},{1,1},{1,3},{2,3},{3,2},{3,1}}

	// result: 
	// row := int(0)
	// col := int(0)
	// cells := [][]int{}

	result := latestDayToCross(row, col, cells)
	fmt.Printf("result = %v\n", result)
}

