package main
import (
	"fmt"
)

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	} else if len(nums) == 1 {
		return nums
	}

	result := make([]int, len(nums))
	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		result[i] = result[i - 1] + num
	}
	return result
}

func main() {
	// result: [1, 3, 6, 10]
	// nums := []int{1, 2, 3, 4}

	// result: [1, 2, 3, 4, 5]
	// nums := []int{1, 1, 1, 1, 1}

	// result: [3, 4, 6, 16, 17]
	nums := []int{3, 1, 2, 10, 1}

	result := runningSum(nums)
	fmt.Printf("result = %v\n", result)
}

