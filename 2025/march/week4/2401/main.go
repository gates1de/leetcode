package main
import (
	"fmt"
)

func longestNiceSubarray(nums []int) int {
	usedBits := int(0)
	windowStart := int(0)
	result := int(0)

	for i := 0; i < len(nums); i++ {
		for (usedBits & nums[i]) != 0 {
			usedBits ^= nums[windowStart]
			windowStart++
		}

		usedBits |= nums[i]
		result = max(result, i - windowStart + 1)
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
	// result: 3
	// nums := []int{1,3,8,48,10}

	// result: 1
	nums := []int{3,1,5,11,13}

	// result: 
	// nums := []int{}

	result := longestNiceSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

