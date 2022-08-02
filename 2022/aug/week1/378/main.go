package main
import (
	"fmt"
	"math"
)

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return -1
	}

	indexMemo := map[int]int{}
	result := int(-1)
	for k > 0 {
		minNum := math.MaxInt32
		minRow := int(0)
		for row := 0; row < n; row++ {
			index := indexMemo[row]
			if index >= n {
				continue
			}

			num := matrix[row][index]
			if num < minNum {
				minNum = num
				minRow = row
			}
		}

		k--
		if k == 0 {
			result = matrix[minRow][indexMemo[minRow]]
			break
		}
		indexMemo[minRow]++
	}

	return result
}

func main() {
	// result: 13
	// matrix := [][]int{{1,5,9},{10,11,13},{12,13,15}}
	// k := int(8)

	// result: -5
	matrix := [][]int{{-5}}
	k := int(1)

	// result: 
	// matrix := [][]int{}
	// k := int(0)

	result := kthSmallest(matrix, k)
	fmt.Printf("result = %v\n", result)
}

