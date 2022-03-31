package main
import (
	"fmt"
)

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if nums[0] > target{
		return helper(-1, 0, len(nums), nums, target)
	}
	return helper(0, 0, len(nums), nums, target)
}

func helper(index int, count int, n int, nums []int, target int) bool {
	if count == n {
		return false
	}

	if index < 0 {
		num := nums[index + n]
		if num == target {
			return true
		}
		return helper(index - 1, count + 1, n, nums, target)
	}

	num := nums[index]
	if num == target {
		return true
	}
	return helper(index + 1, count + 1, n, nums, target)
}

func main() {
	// result: true
	// nums := []int{2,5,6,0,0,1,2}
	// target := int(0)

	// result: false
	// nums := []int{2,5,6,0,0,1,2}
	// target := int(3)

	// result: false
	// nums := []int{1}
	// target := int(0)

	// result: true
	// nums := []int{1,2,3,4,5,6,7,8,9}
	// target := int(5)

	// result: false
	nums := []int{4,5,6,7,8,9,10,-1}
	target := int(-3)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := search(nums, target)
	fmt.Printf("result = %v\n", result)
}

