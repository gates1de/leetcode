package main
import (
	"fmt"
)

func maxSubarrayLength(nums []int, k int) int {
	result := int(0)
	start := int(-1)
	freq := make(map[int]int)

	for end := 0; end < len(nums); end++ {
		freq[nums[end]]++
		for freq[nums[end]] > k {
			start++
			freq[nums[start]]--
		}
		result = max(result, end - start)
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
	// nums := []int{1,2,3,1,2,3,1,2}
	// k := int(2)

	// result: 2
	// nums := []int{1,2,1,2,1,2,1,2}
	// k := int(1)

	// result: 4
	nums := []int{5,5,5,5,5,5,5}
	k := int(4)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxSubarrayLength(nums, k)
	fmt.Printf("result = %v\n", result)
}

