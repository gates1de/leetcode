package main
import (
	"fmt"
)

func findFarmland(land [][]int) [][]int {
	n := len(land)
	m := len(land[0])

	result := make([][]int, 0, n)

	for row1 := 0; row1 < n; row1++ {
		for col1 := 0; col1 < m; col1++ {
			if land[row1][col1] == 1 {
				x := row1
				y := col1

				for x = row1; x < n && land[x][col1] == 1; x++ {
					for y = col1; y < m && land[x][y] == 1; y++ {
						land[x][y] = 0
					}
				}

				result = append(result, []int{row1, col1, x - 1, y - 1})
			}
		}
	}

	return result
}

func main() {
	// result: [[0,0,0,0],[1,1,2,2]]
	// land := [][]int{{1,0,0},{0,1,1},{0,1,1}}

	// result: [[0,0,1,1]]
	// land := [][]int{{1,1},{1,1}}

	// result: []
	land := [][]int{{0}}

	// result: []
	// land := [][]int{}

	result := findFarmland(land)
	fmt.Printf("result = %v\n", result)
}

