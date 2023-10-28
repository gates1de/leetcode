package main
import (
	"fmt"
)

func maximumScore(nums []int, k int) int {
	n := len(nums)
	left := k
	right := k
	result := nums[k]
	currentMin := nums[k]
	for left > 0 || right < n - 1 {
		leftNum := 0
		if left > 0 {
			leftNum = nums[left - 1]
		}

		rightNum := 0
		if right < n - 1 {
			rightNum = nums[right + 1]
		}

		if leftNum < rightNum {
			right++
			currentMin = min(currentMin, nums[right])
		} else {
			left--
			currentMin = min(currentMin, nums[left])
		}

		result = max(result, currentMin * (right - left + 1))
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
	// result: 15
	// nums := []int{1,4,3,7,4,5}
	// k := int(3)

	// result: 20
	nums := []int{5,5,4,5,4,1,1,1}
	k := int(0)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maximumScore(nums, k)
	fmt.Printf("result = %v\n", result)
}

