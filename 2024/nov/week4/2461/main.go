package main
import (
	"fmt"
)

func maximumSubarraySum(nums []int, k int) int64 {
	result := int64(0)
	sum := int64(0)
	begin := int(0)
	end := int(0)
	numToIndex := make(map[int]int)

	for end < len(nums) {
		num := nums[end]
		last := -1

		if i, exists := numToIndex[num]; exists {
			last = i
		}

		for begin <= last || (end - begin + 1 > k) {
			sum -= int64(nums[begin])
			begin++
		}

		numToIndex[num] = end
		sum += int64(nums[end])

		if end - begin + 1 == k {
			result = max(result, sum)
		}

		end++
	}

	return result
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 15
	// nums := []int{1,5,4,2,9,9}
	// k := int(3)

	// result: 0
	nums := []int{4,4,4}
	k := int(3)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maximumSubarraySum(nums, k)
	fmt.Printf("result = %v\n", result)
}

