package main
import (
	"fmt"
)

func countSubarrays(nums []int, k int) int64 {
	maxElement := int(0)
	for _, num := range nums {
		maxElement = max(maxElement, num)
	}

	result := int64(0)
	start := int(0)
	maxElementsInWindow := int(0)

	for end := 0; end < len(nums); end++ {
		if nums[end] == maxElement {
			maxElementsInWindow++
		}

		for k == maxElementsInWindow {
			if nums[start] == maxElement {
				maxElementsInWindow--
			}
			start++
		}

		result += int64(start)
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
	// result: 6
	// nums := []int{1,3,2,3,3}
	// k := int(2)

	// result: 0
	nums := []int{1,4,2,1}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := countSubarrays(nums, k)
	fmt.Printf("result = %v\n", result)
}
