package main
import (
	"fmt"
)

func subsetXORSum(nums []int) int {
	result := int(0)

	for _, num := range nums {
		result |= num
	}

	return result << (len(nums) - 1)
}

func main() {
	// result: 6
	// nums := []int{1,3}

	// result: 28
	// nums := []int{5,1,6}

	// result: 480
	nums := []int{3,4,5,6,7,8}

	// result: 
	// nums := []int{}

	result := subsetXORSum(nums)
	fmt.Printf("result = %v\n", result)
}

