package main
import (
	"fmt"
)

func minOperations(nums []int) int {
	n := len(nums)
	result := int(0)

	for i := 0; i <= n - 3; i++ {
		if nums[i] != 0 {
			continue
		}

		nums[i] = 1
		if nums[i + 1] == 0 {
			nums[i + 1] = 1
		} else {
			nums[i + 1] = 0
		}
		if nums[i + 2] == 0 {
			nums[i + 2] = 1
		} else {
			nums[i + 2] = 0
		}

		result++
	}

	if nums[n - 2] == 0 || nums[n - 1] == 0 {
		return -1
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{0,1,1,1,0,0}

	// result: -1
	nums := []int{0,1,1,1}

	// result: 
	// nums := []int{}

	result := minOperations(nums)
	fmt.Printf("result = %v\n", result)
}

