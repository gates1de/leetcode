package main
import (
	"fmt"
)

func longestSubarray(nums []int) int {
	maxNum := int(0)
	result := int(0)
	currentStreak := int(0)

	for _, num := range nums {
		if maxNum < num {
			maxNum = num
			result = 0
			currentStreak = 0
		}

		if maxNum == num {
			currentStreak++
		} else {
			currentStreak = 0
		}

		result = max(result, currentStreak)
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
	// result: 2
	// nums := []int{1,2,3,3,2,2}

	// result: 1
	nums := []int{1,2,3,4}

	// result: 
	// nums := []int{}

	result := longestSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

