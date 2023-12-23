package main
import (
	"fmt"
	"sort"
)

func maxProductDifference(nums []int) int {
	n := len(nums)
	if n < 4 {
		return -1
	}

	sort.Ints(nums)
	return (nums[n - 2] * nums[n - 1]) - (nums[0] * nums[1])
}

func main() {
	// result: 34
	// nums := []int{5,6,2,7,4}

	// result: 64
	nums := []int{4,2,5,9,7,4,8}

	// result: 
	// nums := []int{}

	result := maxProductDifference(nums)
	fmt.Printf("result = %v\n", result)
}

