package main
import (
	"fmt"
)

func rob(nums []int) int {
	result := int(0)
	helper(0, false, map[int]int{}, nums, &result)
	return result
}

func helper(sum int, isRobbed bool, memo map[int]int, nums[]int, result *int) {
	if len(nums) == 0 {
		if *result < sum {
			*result = sum
		}
		return
	}

	if isRobbed {
		if memo[len(nums)] > sum {
			return
		}
		memo[len(nums)] = sum
	}

	money := nums[0]
	if !isRobbed && money > 0 {
		helper(sum + money, true, memo, nums[1:], result)
	}
	helper(sum, false, memo, nums[1:], result)
}

func main() {
	// result: 4
	// nums := []int{1, 2, 3, 1}

	// result: 12
	// nums := []int{2, 7, 9, 3, 1}

	// result: 1
	// nums := []int{1}

	// result: 0
	// nums := []int{0}

	// result: 123
	// nums := []int{1, 0, 10, 3, 7, 5, 2, 1, 3, 4, 100}

	// result: 4173
	// nums := []int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,205,94,205,169,241,202,144,240}

	// result: 0
	nums := []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

	// result: 
	// nums := []int{}

	result := rob(nums)
	fmt.Printf("result = %v\n", result)
}

