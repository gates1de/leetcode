package main
import (
	"fmt"
)

func continuousSubarrays(nums []int) int64 {
	right := int(0)
	left := int(0)
	currentMin := nums[right]
	currentMax := nums[right]
	windowLen := int(0)
	result := int64(0)

	for right < len(nums) {
		currentMin = min(currentMin, nums[right])
		currentMax = max(currentMax, nums[right])

		if currentMax - currentMin > 2 {
			windowLen = right - left
			result += int64((windowLen * (windowLen + 1)) / 2)

			left = right
			currentMin = nums[right]
			currentMax = nums[right]

			for left > 0 && abs(nums[right] - nums[left - 1]) <= 2 {
				left--
				currentMin = min(currentMin, nums[left])
				currentMax = max(currentMax, nums[left])
			}

			if left < right {
				windowLen = right - left
				result -= int64((windowLen * (windowLen + 1)) / 2)
			}
		}

		right++
	}

	windowLen = right - left
	result += int64((windowLen * (windowLen + 1)) / 2)

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
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
	// result: 8
	// nums := []int{5,4,2,4}

	// result: 6
	nums := []int{1,2,3}

	// result: 
	// nums := []int{}

	result := continuousSubarrays(nums)
	fmt.Printf("result = %v\n", result)
}

