package main
import (
	"fmt"
)

func isZeroArray(nums []int, queries [][]int) bool {
	arr := make([]int, len(nums) + 1)
	for _, query := range queries {
		left := query[0]
		right := query[1]
		arr[left] += 1
		arr[right + 1] -= 1
	}

	counter := make([]int, len(arr))
	operation := int(0)
	for i, delta := range arr {
		operation += delta
		counter[i] = operation
	}

	for i := 0; i < len(nums); i++ {
		if counter[i] < nums[i] {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// nums := []int{1,0,1}
	// queries := [][]int{{0,2}}

	// result: false
	nums := []int{4,3,2,1}
	queries := [][]int{{1,3},{0,2}}

	// result: 
	// nums := []int{}
	// queries := [][]int{}

	result := isZeroArray(nums, queries)
	fmt.Printf("result = %v\n", result)
}

