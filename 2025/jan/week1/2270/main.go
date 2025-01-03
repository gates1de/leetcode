package main
import (
	"fmt"
)

func waysToSplitArray(nums []int) int {
	leftSum := int(0)
	rightSum := int(0)

	for _, num := range nums {
		rightSum += num
	}

	result := int(0)

	for i := 0; i < len(nums) - 1; i++ {
		leftSum += nums[i]
		rightSum -= nums[i]

		if leftSum >= rightSum {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{10,4,-8,7}

	// result: 2
	nums := []int{2,3,1,0}

	// result: 
	// nums := []int{}

	result := waysToSplitArray(nums)
	fmt.Printf("result = %v\n", result)
}

