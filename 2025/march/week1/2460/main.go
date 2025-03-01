package main
import (
	"fmt"
)

func applyOperations(nums []int) []int {
	n := len(nums)

	for i := 0; i < n - 1; i++ {
		if nums[i] == nums[i + 1] && nums[i] != 0 {
			nums[i] *= 2
			nums[i + 1] = 0
		}
	}

	nonZeroIndex := int(0)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for nonZeroIndex < n {
		nums[nonZeroIndex] = 0
		nonZeroIndex++
	}

	return nums
}

func main() {
	// result: [1,4,2,0,0,0]
	// nums := []int{1,2,2,1,1,0}

	// result: [1,0]
	nums := []int{0,1}

	// result: []
	// nums := []int{}

	result := applyOperations(nums)
	fmt.Printf("result = %v\n", result)
}

