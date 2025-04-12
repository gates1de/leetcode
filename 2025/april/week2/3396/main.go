package main
import (
	"fmt"
)

func minimumOperations(nums []int) int {
	seen := make([]bool, 128)

	for i := len(nums) - 1; i >= 0; i-- {
		if seen[nums[i]] {
			return i / 3 + 1
		}
		seen[nums[i]] = true
	}

	return 0
}

func main() {
	// result: 3
	// nums := []int{1,2,3,4,2,3,3,5,7}

	// result: 2
	// nums := []int{4,5,6,4,4}

	// result: 0
	nums := []int{6,7,8,9}

	// result: 
	// nums := []int{}

	result := minimumOperations(nums)
	fmt.Printf("result = %v\n", result)
}

