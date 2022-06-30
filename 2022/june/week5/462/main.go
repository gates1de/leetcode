package main
import (
	"fmt"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	med := nums[n / 2]
	result := diffSum(med, nums)
	if n % 2 == 0 {
		med = nums[n / 2 - 1]
		result = min(result, diffSum(med, nums))
	}
	return result
}

func diffSum(target int, nums []int) int {
	result := int(0)
	for _, num := range nums {
		if target < num {
			result += num - target
		} else {
			result += target - num
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// nums := []int{1, 2, 3}

	// result: 16
	// nums := []int{1, 10, 2, 9}

	// result: 
	nums := []int{3,-20,8,14,-6,2,30,4,2,6,10,-4,-6,7,2,-30,1,8,46,8,3,0}

	// result: 
	// nums := []int{}

	result := minMoves2(nums)
	fmt.Printf("result = %v\n", result)
}

