package main
import (
	"fmt"
)

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	sum := int(0)
	result := int(0)
	diff := make([]int, n + 1)

	for i, num := range nums {
		for sum + diff[i] < num {
			result++

			if result > len(queries) {
				return -1
			}

			query := queries[result - 1]
			left := query[0]
			right := query[1]
			val := query[2]

			if right >= i {
				diff[max(left, i)] += val
				diff[right + 1] -= val
			}
		}

		sum += diff[i]
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// nums := []int{2,0,2}
	// queries := [][]int{{0,2,1},{0,2,1},{1,1,3}}

	// result: -1
	nums := []int{4,3,2,1}
	queries := [][]int{{1,3,2},{0,2,1}}

	// result: 
	// nums := []int{}
	// queries := [][]int{}

	result := minZeroArray(nums, queries)
	fmt.Printf("result = %v\n", result)
}

