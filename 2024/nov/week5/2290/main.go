package main
import (
	"fmt"
)

func minimumObstacles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i, _ := range seen {
		seen[i] = make([]bool, n)
	}

	current := [][2]int{{0, 0}}
	next := [][2]int{}
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	seen[0][0] = true
	for steps := 0; len(current) > 0; steps++ {
		next = next[:0]

		for k := 0; k < len(current); k++ {
			i := current[k][0]
			j := current[k][1]

			for _, dir := range dirs {
				ii := i + dir[0]
				jj := j + dir[1]

				if !isValid(grid, ii, jj) || seen[ii][jj] {
					continue
				}

				if ii == m - 1 && jj == n - 1 {
					return steps
				}

				seen[ii][jj] = true

				if grid[ii][jj] == 1 {
					next = append(next, [2]int{ii, jj})
				} else {
					current = append(current, [2]int{ii, jj})
				}
			}
		}

		current, next = next, current
	}

	return -1
}

func isValid(grid [][]int, row int, col int) bool {
	return row >= 0 && col >= 0 && row < len(grid) && col < len(grid[0])
}

func main() {
	// result: 2
	// grid := [][]int{{0,1,1},{1,1,0},{1,1,0}}

	// result: 0
	grid := [][]int{{0,1,0,0,0},{0,1,0,1,0},{0,0,0,1,0}}

	// result: 
	// grid := [][]int{}

	result := minimumObstacles(grid)
	fmt.Printf("result = %v\n", result)
}

