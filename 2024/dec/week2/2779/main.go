package main
import (
	"fmt"
)

func maximumBeauty(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	result := int(0)
	maxVal := int(0)

	for _, num := range nums {
		maxVal = max(maxVal, num)
	}

	counter := make([]int, maxVal + 1)

	for _, num := range nums {
		counter[max(num - k, 0)]++
		counter[min(num + k + 1, maxVal)]--
	}

	currentSum := int(0)
	for _, val := range counter {
		currentSum += val
		result = max(result, currentSum)
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
	// result: 3
	// nums := []int{4,6,1,2}
	// k := int(2)

	// result: 4
	nums := []int{1,1,1,1}
	k := int(10)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maximumBeauty(nums, k)
	fmt.Printf("result = %v\n", result)
}

