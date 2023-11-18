package main
import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left := int(0)
	result := int(0)
	curr := int64(0)

	for right := 0; right < len(nums); right++ {
		target := nums[right]
		curr += int64(target)

		for int64((right - left + 1) * target) - curr > int64(k) {
			curr -= int64(nums[left])
			left++
		}

		result = max(result, right - left + 1)
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
	// result: 3
	// nums := []int{1,2,4}
	// k := int(5)

	// result: 2
	// nums := []int{1,4,8,13}
	// k := int(5)

	// result: 1
	nums := []int{3,9,6}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxFrequency(nums, k)
	fmt.Printf("result = %v\n", result)
}

