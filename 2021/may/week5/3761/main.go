package main
import (
	"fmt"
	"sort"
)

func maximumGap(nums []int) int {
	result := int(0)
	sort.Slice(nums, func(a, b int) bool { return nums[a] < nums[b] })
	for i := 0; i < len(nums) - 1; i++ {
		current := nums[i]
		next := nums[i + 1]
		diff := next - current
		if result < diff {
			result = diff
		}
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{3, 6, 9, 1}

	// result: 0
	// nums := []int{10}

	// result: 0
	// nums := []int{}

	// result: 10
	nums := []int{10, 20}

	// result: 
	// nums := []int{}

	// result: 
	// nums := []int{}

	result := maximumGap(nums)
	fmt.Printf("result = %v\n", result)
}

