package main
import (
	"fmt"
	"math"
)

func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	result := math.MinInt32
	for i := 1; i < n; i++ {
		result = max(result, diff(nums[i - 1], nums[i]))
	}

	result = max(result, diff(nums[n - 1], nums[0]))
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func diff(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: 3
	// nums := []int{1,2,4}

	// result: 5
	nums := []int{-5,-10,-5}

	// result: 
	// nums := []int{}

	result := maxAdjacentDistance(nums)
	fmt.Printf("result = %v\n", result)
}

