package main
import (
	"fmt"
)

func getMaximumXor(nums []int, maximumBit int) []int {
	xor := int(0)

	for _, num := range nums {
		xor ^= num
	}

	result := make([]int, len(nums))
	mask := (1 << maximumBit) - 1
	for i, _ := range nums {
		result[i] = xor ^ mask
		xor ^= nums[len(nums) - 1 - i]
	}

	return result
}

func main() {
	// result: [0,3,2,3]
	// nums := []int{0,1,1,3}
	// maximumBit := int(2)

	// result: [5,2,6,5]
	// nums := []int{2,3,4,7}
	// maximumBit := int(3)

	// result: [4,3,6,4,6,7]
	nums := []int{0,1,2,2,5,7}
	maximumBit := int(3)

	// result: []
	// nums := []int{}
	// maximumBit := int(0)

	result := getMaximumXor(nums, maximumBit)
	fmt.Printf("result = %v\n", result)
}

