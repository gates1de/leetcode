package main
import (
	"fmt"
)

func countSubarrays(nums []int) int {
	n := len(nums)
	result := int(0)

	for i := 0; i < n - 2; i++ {
		if (nums[i] + nums[i + 2]) * 2 == nums[i + 1] {
			result++
		}
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{1,2,1,4,1}

	// result: 0
	// nums := []int{1,1,1}

	// result: 1
	nums := []int{-1,-4,-1,4}

	// result: 
	// nums := []int{}

	result := countSubarrays(nums)
	fmt.Printf("result = %v\n", result)
}

