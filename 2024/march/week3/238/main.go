package main
import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	product := int(0)
	zeroCount := int(0)

	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}

		if product == 0 {
			product = num
			continue
		}

		product *= num
	}

	if zeroCount >= 2 {
		return make([]int, len(nums))
	}

	for i, num := range nums {
		if num == 0 {
			nums[i] = product
			continue
		}

		if zeroCount > 0 {
			nums[i] = 0
			continue
		}

		nums[i] = product / num
	}

	return nums
}

func main() {
	// result: [24,12,8,6]
	// nums := []int{1,2,3,4}

	// result: [0,0,9,0,0]
	nums := []int{-1,1,0,-3,3}

	// result: []
	// nums := []int{}

	result := productExceptSelf(nums)
	fmt.Printf("result = %v\n", result)
}

