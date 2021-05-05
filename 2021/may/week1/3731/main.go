package main
import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	isModified := false
	for i := 0; i < len(nums) - 1; i++ {
		// fmt.Printf("nums = %v, nums[i] = %v, nums[i + 1] = %v\n", nums, nums[i], nums[i + 1])
		if nums[i] > nums[i + 1] {
			if isModified {
				return false
			}
			isModified = true
			if i == 0 {
				nums[i] = nums[i + 1]
			} else {
				if nums[i - 1] > nums[i + 1] {
					nums[i + 1] = nums[i]
				} else {
					nums[i] = nums[i - 1]
				}
			}
		}
	}
	if isModified && nums[len(nums) - 2] > nums[len(nums) - 1] {
		return false
	}
	return true
}

func main() {
	// result: true
	// nums := []int{4, 2, 3}

	// result: false
	// nums := []int{4, 2, 1}

	// result: false
	// nums := []int{3, 4, 2, 3}

	// result: true
	nums := []int{5, 7, 1, 8}

	// result: 
	// nums := []int{}

	result := checkPossibility(nums)
	fmt.Printf("result = %v\n", result)
}

