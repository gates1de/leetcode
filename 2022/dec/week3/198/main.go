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
	// nums := []int{1,2,3,1}

	// result: 12
	nums := []int{2,7,9,3,1}

	// result: 
	// nums := []int{}

	result := rob(nums)
	fmt.Printf("result = %v\n", result)
}

