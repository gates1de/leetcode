package main
import (
	"fmt"
)

func minimumReplacement(nums []int) int64 {
	n := len(nums)
	result := int64(0)

	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i + 1] {
			continue
		}

		numElements := int64(nums[i] + nums[i + 1] - 1) / int64(nums[i + 1])
		result += numElements - 1
		nums[i] /= int(numElements)
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{3,9,3}

	// result: 0
	nums := []int{1,2,3,4,5}

	// result: 
	// nums := []int{}

	result := minimumReplacement(nums)
	fmt.Printf("result = %v\n", result)
}

