package main
import (
	"fmt"
)

func maxSubarraySumCircular(nums []int) int {
	currentMax := int(0)
	currentMin := int(0)
	sum := int(0)
	maxSum := nums[0]
	minSum := nums[0]
	for _, num := range nums {
		currentMax = max(currentMax, 0) + num
		maxSum = max(maxSum, currentMax)
		currentMin = min(currentMin, 0) + num
		minSum = min(minSum, currentMin)
		sum += num
	}

	if sum == minSum {
		return maxSum
	}
	return max(maxSum, sum - minSum)
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
	// result: 3
	// nums := []int{1,-2,3,-2}

	// result: 10
	// nums := []int{5,-3,5}

	// result: -2
	nums := []int{-3,-2,-3}

	// result: 
	// nums := []int{}

	result := maxSubarraySumCircular(nums)
	fmt.Printf("result = %v\n", result)
}

