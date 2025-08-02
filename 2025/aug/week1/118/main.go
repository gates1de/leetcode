package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
    result := make([][]int, numRows)
    if numRows == 0 {
        return result
    }

    result[0] = []int{1}

    for i := 1; i < numRows; i++ {
        result[i] = make([]int, i + 1)
        for j := 0; j <= i; j++ {
            if j - 1 >= 0 {
                result[i][j] += result[i - 1][j - 1]
            }
            if j < i {
                result[i][j] += result[i - 1][j]
            }
        }
    }

    return result
}

func main() {
	// result: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	// numRows := int(5)

	// result: [[1]]
	numRows := int(1)

	// result: 
	// numRows := int(0)

	result := generate(numRows)
	fmt.Printf("result = %v\n", result)
}

