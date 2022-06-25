package main
import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums) - k]
}

func main() {
	// result: 5
	// nums := []int{3,2,1,5,6,4}
	// k := int(2)

	// result: 4
	nums := []int{3,2,3,1,2,4,5,5,6}
	k := int(4)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := findKthLargest(nums, k)
	fmt.Printf("result = %v\n", result)
}

