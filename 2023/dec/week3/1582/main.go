package main
import (
	"fmt"
)

func numSpecial(mat [][]int) int {
	m := len(mat)
	if m == 0 {
		return 0
	}

	n := len(mat[0])
	if n == 0 {
		return 0
	}

	rowCount := make([]int, m)
	colCount := make([]int, n)

	for row, _ := range mat {
		for col, val := range mat[row] {
			if val != 1 {
				continue
			}
			rowCount[row]++
			colCount[col]++
		}
	}

	result := int(0)
	for row, _ := range mat {
		for col, val := range mat[row] {
			if val != 1 || rowCount[row] != 1 || colCount[col] != 1 {
				continue
			}

			result++
		}
	}

	return result
}

func main() {
	// result: 1
	// mat := [][]int{{1,0,0},{0,0,1},{1,0,0}}

	// result: 3
	mat := [][]int{{1,0,0},{0,1,0},{0,0,1}}

	// result: 
	// mat := [][]int{}

	result := numSpecial(mat)
	fmt.Printf("result = %v\n", result)
}

