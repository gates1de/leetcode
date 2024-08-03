package main
import (
	"fmt"
)

func minSwaps(nums []int) int {
	return min(helper(nums, 0), helper(nums, 1))
}

func helper(nums []int, val int) int {
	n := len(nums)
	count := int(0)

	for i := n - 1; i >= 0; i-- {
		if nums[i] == val {
			count++
		}
	}

	if count == 0 || count == n {
		return 0
	}

	start := int(0)
	end := int(0)
	maxValInWindow := int(0)
	currentValInWindow := int(0)

	for end < count {
		if nums[end] == val {
			currentValInWindow++
		}
		end++
	}

	maxValInWindow = max(maxValInWindow, currentValInWindow)

	for end < n {
		if nums[start] == val {
			currentValInWindow--
		}
		if nums[end] == val {
			currentValInWindow++
		}
		start++
		end++
		maxValInWindow = max(maxValInWindow, currentValInWindow)
	}

	return count - maxValInWindow
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
	// result: 1
	// nums := []int{0,1,0,1,1,0,0}

	// result: 2
	// nums := []int{0,1,1,1,0,0,1,1,0}

	// result: 0
	nums := []int{1,1,0,0,1}

	// result: 
	// nums := []int{}

	result := minSwaps(nums)
	fmt.Printf("result = %v\n", result)
}

