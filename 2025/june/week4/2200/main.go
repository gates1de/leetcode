package main
import (
	"fmt"
)

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	tmpIndex := int(0)
	result := make([]int, 0)

	for j := 0; j < n; j++ {
		if nums[j] == key {
			l := max(tmpIndex, j - k)
			tmpIndex = min(n - 1, j + k) + 1

			for i := l; i < tmpIndex; i++ {
				result = append(result, i)
			}
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [1,2,3,4,5,6]
	// nums := []int{3,4,9,1,3,9,5}
	// key := int(9)
	// k := int(1)

	// result: [0,1,2,3,4]
	nums := []int{2,2,2,2,2}
	key := int(2)
	k := int(2)

	// result: []
	// nums := []int{}
	// key := int(0)
	// k := int(0)

	result := findKDistantIndices(nums, key, k)
	fmt.Printf("result = %v\n", result)
}

