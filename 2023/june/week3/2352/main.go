package main
import (
	"fmt"
)

func equalPairs(grid [][]int) int {
	n := len(grid)
	cols := map[string]int{}
	rows := map[string]int{}

	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			str += fmt.Sprintf("%v,", grid[i][j])
		}
		rows[str]++
	}

	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n; j++ {
			str += fmt.Sprintf("%v,", grid[j][i])
		}
		cols[str]++
	}

	result := int(0)
	for col, colCount := range cols {
		for row, rowCount := range rows {
			if col == row {
				result += colCount * rowCount
			}
		}
	}

	return result
}

func main() {
	// result: 1
	// grid := [][]int{{3,2,1},{1,7,6},{2,7,7}}

	// result: 3
	grid := [][]int{{3,1,2,2},{1,4,4,5},{2,4,2,2},{2,4,2,2}}

	// result: 
	// grid := [][]int{}

	result := equalPairs(grid)
	fmt.Printf("result = %v\n", result)
}

