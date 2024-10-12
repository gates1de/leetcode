package main
import (
	"fmt"
)

func maxWidthRamp(nums []int) int {
	n := len(nums)
	stack := make([]int, 0, n)

	for i, num := range nums {
		if len(stack) == 0 || nums[stack[len(stack) - 1]] > num {
			stack = append(stack, i)
		}
	}

	result := int(0)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack) - 1]] <= nums[i] {
			result = max(result, i - stack[len(stack) - 1])
			stack = stack[:len(stack) - 1]
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

func main() {
	// result: 4
	// nums := []int{6,0,8,2,1,5}

	// result: 7
	nums := []int{9,8,1,0,1,9,4,0,4,1}

	// result: 
	// nums := []int{}

	result := maxWidthRamp(nums)
	fmt.Printf("result = %v\n", result)
}

