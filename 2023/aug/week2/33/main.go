package main
import (
	"fmt"
)

func search(nums []int, target int) int {
	n := len(nums)
	left := int(0)
	right := n - 1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	// result: 4
	// nums := []int{4,5,6,7,0,1,2}
	// target := int(0)

	// result: -1
	// nums := []int{4,5,6,7,0,1,2}
	// target := int(-4)

	// result: -1
	nums := []int{1}
	target := int(0)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := search(nums, target)
	fmt.Printf("result = %v\n", result)
}

