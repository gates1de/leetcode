package main
import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	n := len(nums)
	startIndex := int(0)
	endIndex := n - 1

	for startIndex <= endIndex {
		midIndex := (startIndex + endIndex) / 2
		mid := nums[midIndex]
		if mid == target {
			return midIndex
		} else if mid > target {
			endIndex = midIndex - 1
		} else {
			startIndex = midIndex + 1
		}
	}

	return endIndex + 1
}

func main() {
	// result: 2
	// nums := []int{1,3,5,6}
	// target := int(5)

	// result: 1
	// nums := []int{1,3,5,6}
	// target := int(2)

	// result: 4
	nums := []int{1,3,5,6}
	target := int(7)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := searchInsert(nums, target)
	fmt.Printf("result = %v\n", result)
}

