package main
import (
	"fmt"
)

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	maxReach := make([]int, n)
	end := int(0)

	for i := 0; i < n; i++ {
		end = max(end, i)
		for end < n - 1 && nums[end] % 2 != nums[end + 1] % 2 {
			end++
		}

		maxReach[i] = end
	}

	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = q[1] <= maxReach[q[0]]
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
	// result: [false]
	// nums := []int{3,4,1,2,6}
	// queries := [][]int{{0,4}}

	// result: [false,true]
	nums := []int{4,3,1,6}
	queries := [][]int{{0,2},{2,3}}

	// result: []
	// nums := []int{}
	// queries := [][]int{}

	result := isArraySpecial(nums, queries)
	fmt.Printf("result = %v\n", result)
}

