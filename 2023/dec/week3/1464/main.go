package main
import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	sort.Ints(nums)
	return (nums[n - 1] - 1) * (nums[n - 2] - 1)
}

func main() {
	// result: 12
	// nums := []int{3,4,5,2}

	// result: 16
	// nums := []int{1,5,4,5}

	// result: 12
	nums := []int{3,7}

	// result: 
	// nums := []int{}

	result := maxProduct(nums)
	fmt.Printf("result = %v\n", result)
}

