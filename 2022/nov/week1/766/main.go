package main
import (
	"fmt"
)

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])

	for i := m - 1; i >= 0; i-- {
		y := i
		x := int(0)
		pre := int(-1)

		for y < m && x < n {
			if pre != -1 && pre != matrix[y][x] {
				return false
			}
			pre = matrix[y][x]
			y++
			x++
		}
	}

	for i := 1; i < n; i++ {
		y := int(0)
		x := i
		pre := int(-1)

		for y < m && x < n {
			if pre != -1 && pre != matrix[y][x] {
				return false
			}
			pre = matrix[y][x]
			y++
			x++
		}
	}

	return true
}

func main() {
	// result: true
	// matrix := [][]int{{1,2,3,4},{5,1,2,3},{9,5,1,2}}

	// result: false
	// matrix := [][]int{{1,2},{2,2}}

	// result: false
	matrix := [][]int{{11,74,0,93},{40,11,74,7}}

	// result: 
	// matrix := [][]int{}

	result := isToeplitzMatrix(matrix)
	fmt.Printf("result = %v\n", result)
}

