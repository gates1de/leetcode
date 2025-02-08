package main
import (
	"fmt"
)

func longestMonotonicSubarray(nums []int) int {
	increase := int(1)
	decrease := int(1)
	result := int(1)

	for i := 0; i < len(nums) - 1; i++ {
		if nums[i + 1] > nums[i] {
			increase++
			decrease = 1
		} else if nums[i + 1] < nums[i] {
			decrease++
			increase = 1
		} else {
			increase = 1
			decrease = 1
		}

		result = max(result, max(increase, decrease))
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
	// nums := []int{1,4,3,3,2}

	// result: 1
	// nums := []int{3,3,3,3}

	// result: 3
	nums := []int{3,2,1}

	// result: 
	// nums := []int{}

	result := longestMonotonicSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

