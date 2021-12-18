package main
import (
	"fmt"
)

func canPartition(nums []int) bool {
	sum := int(0)
	for _, num := range nums {
		sum += num
	}

	if sum % 2 != 0 {
		return false
	}

	return helper(0, sum / 2, map[int]bool{}, nums)
}

func helper(current int, target int, m map[int]bool, nums []int) bool {
	if current == target {
		return true
	}
	if m[current] || len(nums) == 0 || current > target {
		m[current] = true
		return false
	}

	if current + nums[0] <= target {
		isTarget := helper(current + nums[0], target, m, nums[1:])
		if isTarget {
			return true
		}
	}
	return helper(current, target, m, nums[1:])
}

func main() {
	// result: true
	// nums := []int{1, 5, 11, 5}

	// result: false
	// nums := []int{1, 2, 3, 5}

	// result: false
	// nums := []int{1}

	// result: true
	// nums := []int{11, 1, 7, 4, 5, 6, 12}

	// result: false
	nums := []int{100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,99,97}

	// result: 
	// nums := []int{}

	result := canPartition(nums)
	fmt.Printf("result = %v\n", result)
}

