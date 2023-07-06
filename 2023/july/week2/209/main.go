package main
import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	left := int(0)
	sumOfCurrentWindow := int(0)
	result := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sumOfCurrentWindow += nums[right]

		for sumOfCurrentWindow >= target {
			result = min(result, right - left + 1)
			sumOfCurrentWindow -= nums[left]
			left++
		}
	}

	if result == math.MaxInt32 {
		return 0
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
	// result: 2
	// target := int(7)
	// nums := []int{2,3,1,2,4,3}

	// result: 1
	// target := int(4)
	// nums := []int{1,4,4}

	// result: 0
	target := int(11)
	nums := []int{1,1,1,1,1,1,1,1}

	// result: 
	// target := int(0)
	// nums := []int{}

	result := minSubArrayLen(target, nums)
	fmt.Printf("result = %v\n", result)
}

