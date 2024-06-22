package main
import (
	"fmt"
)

func numberOfSubarrays(nums []int, k int) int {
	return helper(nums, k) - helper(nums, k - 1)
}

func helper(nums []int, k int) int {
	window := int(0)
	start := int(0)
	result := int(0)

	for end := 0; end < len(nums); end++ {
		window += nums[end] % 2
		for window > k {
			window -= nums[start] % 2
			start++
		}

		result += end - start + 1
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{1,1,2,1,1}
	// k := int(3)

	// result: 0
	// nums := []int{2,4,6j}
	// k := int(1)

	// result: 16
	nums := []int{2,2,2,1,2,2,1,2,2,2}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := numberOfSubarrays(nums, k)
	fmt.Printf("result = %v\n", result)
}

