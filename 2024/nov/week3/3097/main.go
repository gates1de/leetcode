package main
import (
	"fmt"
	"math"
)

func minimumSubarrayLength(nums []int, k int) int {
	result := math.MaxInt32
	windowStart := int(0)
	windowEnd := int(0)
	bitCounts := make([]int, 32)

	for windowEnd < len(nums) {
		updateBitCounts(bitCounts, nums[windowEnd], 1)

		for windowStart <= windowEnd && convertBitCountsToNumber(bitCounts) >= k {
			result = min(result, windowEnd - windowStart + 1)
			updateBitCounts(bitCounts, nums[windowStart], -1)
			windowStart++
		}

		windowEnd++
	}

	if result == math.MaxInt32 {
		return -1
	}

	return result
}

func updateBitCounts(bitCounts []int, num int, delta int) {
	for bitPosition := 0; bitPosition < 32; bitPosition++ {
		if ((num >> bitPosition) & 1) != 0 {
			bitCounts[bitPosition] += delta
		}
	}
}

func convertBitCountsToNumber(bitCounts []int) int {
	result := int(0)
	for bitPosition := 0; bitPosition < 32; bitPosition++ {
		if bitCounts[bitPosition] != 0 {
			result |= (1 << bitPosition)
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// nums := []int{1,2,3}
	// k := int(2)

	// result: 3
	// nums := []int{2,1,8}
	// k := int(10)

	// result: 1
	nums := []int{1,2}
	k := int(0)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minimumSubarrayLength(nums, k)
	fmt.Printf("result = %v\n", result)
}

