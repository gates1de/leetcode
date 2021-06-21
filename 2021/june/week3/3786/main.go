package main
import (
	"fmt"
)

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		if i == 1 {
			rows := helper(i, []int{})
			result[i - 1] = rows
			continue
		}
		rows := helper(i, result[i - 2])
		result[i - 1] = rows
	}
	return result
}

func helper(level int, preLevelRows []int) []int {
	rows := make([]int, level)
	rows[0] = 1
	if level == 1 {
		return rows
	}
	rows[len(rows) - 1] = 1
	for i := 1; i < len(rows) - 1; i++ {
		rows[i] = preLevelRows[i - 1] + preLevelRows[i]
	}
	return rows
}

func main() {
	// result: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	// numRows := int(5)

	// result: [[1]]
	// numRows := int(1)

	// result: [[1],[1, 1]]
	numRows := int(2)

	// result: 
	// numRows := int(0)

	result := generate(numRows)
	fmt.Printf("result = %v\n", result)
}

