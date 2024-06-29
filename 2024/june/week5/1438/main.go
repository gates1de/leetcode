package main
import (
	"fmt"
)

func longestSubarray(nums []int, limit int) int {
	maxDequeue := make([]int, 0)
	minDequeue := make([]int, 0)
	left := int(0)
	result := int(0)

	for right := 0; right < len(nums); right++ {
		for len(maxDequeue) > 0 && maxDequeue[len(maxDequeue) - 1] < nums[right] {
			maxDequeue = maxDequeue[:len(maxDequeue) - 1]
		}
		maxDequeue = append(maxDequeue, nums[right])

		for len(minDequeue) > 0 && minDequeue[len(minDequeue) - 1] > nums[right] {
			minDequeue = minDequeue[:len(minDequeue) - 1]
		}
		minDequeue = append(minDequeue, nums[right])

		for maxDequeue[0] - minDequeue[0] > limit {
			if maxDequeue[0] == nums[left] {
				maxDequeue = maxDequeue[1:]
			}
			if minDequeue[0] == nums[left] {
				minDequeue = minDequeue[1:]
			}
			left++
		}

		result = max(result, right - left + 1)
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
	// nums := []int{8,2,4,7}
	// limit := int(4)

	// result: 4
	// nums := []int{10,1,2,4,7,2}
	// limit := int(5)

	// result: 3
	nums := []int{4,2,2,2,4,4,2,2}
	limit := int(0)

	// result: 
	// nums := []int{}
	// limit := int(0)

	result := longestSubarray(nums, limit)
	fmt.Printf("result = %v\n", result)
}

