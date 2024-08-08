package main
import (
	"fmt"
)

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	directions := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	result := make([][]int, rows * cols)
	for i, _ := range result {
		result[i] = make([]int, 2)
	}

	index := int(0)
	step := int(1)
	dir := int(0)
	for index < rows * cols {
		for i := 0; i < 2; i++ {
			for j := 0; j < step; j++ {
				if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
					result[index][0] = rStart
					result[index][1] = cStart
					index++
				}

				rStart += directions[dir][0]
				cStart += directions[dir][1]
			}

			dir = (dir + 1) % 4
		}

		step++
	}

	return result
}

func main() {
	// result: [[0,0],[0,1],[0,2],[0,3]]
	// rows := int(1)
	// cols := int(4)
	// rStart := int(0)
	// cStart := int(0)

	// result: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]
	rows := int(5)
	cols := int(6)
	rStart := int(1)
	cStart := int(4)

	// result: []
	// rows := int(0)
	// cols := int(0)
	// rStart := int(0)
	// cStart := int(0)

	result := spiralMatrixIII(rows, cols, rStart, cStart)
	fmt.Printf("result = %v\n", result)
}

