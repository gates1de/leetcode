package main
import (
	"fmt"
)

func maximumDifference(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}

	result := int(-1)
	minNum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > minNum {
			result = max(result, nums[i] - minNum)
		} else {
			minNum = nums[i]
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// nums := []int{7,1,5,4}

	// result: -1
	// nums := []int{9,4,3,2}

	// result: 9
	nums := []int{1,5,2,10}

	// result: 
	// nums := []int{}

	result := maximumDifference(nums)
	fmt.Printf("result = %v\n", result)
}

