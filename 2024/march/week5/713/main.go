package main
import (
	"fmt"
)

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	result := int(0)
	product := int(1)

	for left, right := 0, 0; right < len(nums); right++ {
		product *= nums[right]

		for product >= k {
			product /= nums[left]
			left++
		}

		result += right - left + 1
	}

	return result
}

func main() {
	// result: 8
	// nums := []int{10,5,2,6}
	// k := int(100)

	// result: 0
	nums := []int{1,2,3}
	k := int(0)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := numSubarrayProductLessThanK(nums, k)
	fmt.Printf("result = %v\n", result)
}

