package main
import (
	"fmt"
)

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m == r && n == c || m * n != r * c {
		return mat
	}

	result := make([][]int, r)
	for i, _ := range result {
		result[i] = make([]int, c)
	}

	for i, row := range mat {
		for j, num := range row {
			seq := i * n + j
			targetRow := seq / c
			targetCol := seq % c
			// fmt.Printf("%v, %v: seq = %v, targetRow = %v, targetCol = %v\n", i, j, seq, targetRow, targetCol)
			result[targetRow][targetCol] = num
		}
	}
	return result
}

func main() {
	// result: [[1,2,3,4]]
	// mat := [][]int{{1, 2}, {3, 4}}
	// r := int(1)
	// c := int(4)

	// result: [[1,2],[3,4]]
	// mat := [][]int{{1, 2}, {3, 4}}
	// r := int(2)
	// c := int(4)

	// result: [[0,1],[2,9],[3,4] [5,10] [6,7],[8,11]]
	mat := [][]int{{0, 1, 2, 9}, {3, 4, 5, 10}, {6, 7, 8, 11}}
	r := int(6)
	c := int(2)

	// result: 
	// mat := [][]int{}
	// r := int(0)
	// c := int(0)

	result := matrixReshape(mat, r, c)
	fmt.Printf("result = %v\n", result)
}

