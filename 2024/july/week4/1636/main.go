package main
import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	count := make(map[int]int, len(nums))
	for _, num := range nums {
		count[num]++
	}

	sort.Slice(nums, func (a, b int) bool {
		if count[nums[a]] < count[nums[b]] {
			return true
		} else if count[nums[a]] > count[nums[b]] {
			return false
		}

		return nums[a] > nums[b]
	})

	return nums
}

func main() {
	// result: [3,1,1,2,2,2]
	// nums := []int{1,1,2,2,2,3}

	// result: [1,3,3,2,2]
	// nums := []int{2,3,1,3,2}

	// result: [5,-1,4,4,-6,-6,1,1,1]
	nums := []int{-1,1,-6,4,5,-6,1,4,1}

	// result: []
	// nums := []int{}

	result := frequencySort(nums)
	fmt.Printf("result = %v\n", result)
}

