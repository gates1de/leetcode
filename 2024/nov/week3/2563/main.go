package main
import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	return lowerBound(nums, upper + 1) - lowerBound(nums, lower)
}

func lowerBound(nums []int, val int) int64 {
	left := int(0)
	right := len(nums) - 1
	result := int64(0)

	for left < right {
		sum := nums[left] + nums[right]

		if sum < val {
			result += int64(right - left)
			left++
		} else {
			right--
		}
	}

	return result
}

func main() {
	// result: 6
	// nums := []int{0,1,7,4,4,5}
	// lower := int(3)
	// upper := int(6)

	// result: 1
	nums := []int{1,7,9,2,5}
	lower := int(11)
	upper := int(11)

	// result: 
	// nums := []int{}
	// lower := int(0)
	// upper := int(0)

	result := countFairPairs(nums, lower, upper)
	fmt.Printf("result = %v\n", result)
}

