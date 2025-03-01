package main
import (
	"fmt"
)

func maxAbsoluteSum(nums []int) int {
	minPrefixSum := int(0)
	maxPrefixSum := int(0)
	prefixSum := int(0)

	for _, num := range nums {
		prefixSum += num
		minPrefixSum = min(minPrefixSum, prefixSum)
		maxPrefixSum = max(maxPrefixSum, prefixSum)
	}

	return maxPrefixSum - minPrefixSum
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
	// result: 5
	// nums := []int{1,-3,2,3,-4}

	// result: 8
	nums := []int{2,-5,1,-4,3,-2}

	// result: 
	// nums := []int{}

	result := maxAbsoluteSum(nums)
	fmt.Printf("result = %v\n", result)
}

