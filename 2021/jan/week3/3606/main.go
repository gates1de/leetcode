package main
import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	// fmt.Printf("sortedNums = %v\n", nums)
	return nums[k - 1]
}

func index(nums []int, target int) int {
	for i, val := range nums {
		if val == target {
			return i
		}
	}
	return -1
}

func main() {
	// nums := []int{3, 2, 1, 5, 6, 4}
	// k := int(2)

	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := int(4)

	result := findKthLargest(nums, k)
	fmt.Printf("result = %v\n", result)
}

