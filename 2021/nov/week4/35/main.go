package main
import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left := int(0)
	right := len(nums)
	for left < right {
		middle := (left + right) / 2
		if nums[middle] <= target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	if left > 0 && nums[left - 1] == target {
		return left - 1
	} else {
		return right
	}
}

// Slow
func mySolution(nums []int, target int) int {
	head := int(0)
	tail := len(nums) - 1
	for head <= tail {
		mid := head + (tail - head) / 2
		// fmt.Printf("head = %v, tail = %v, mid = %v, nums = %v, nums[mid] = %v\n", head, tail, mid, nums[head:tail + 1], nums[mid])

		if len(nums[head:tail + 1]) == 1 {
			if nums[mid] >= target {
				return mid
			} else if nums[mid] < target {
				return mid + 1
			}
		}

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			head = mid + 1
		} else if target < nums[mid] {
			tail = mid - 1
		}
	}

	if tail < head {
		return tail + 1
	}

	return head
}

func main() {
	// result: 2
	// nums := []int{1, 3, 5, 6}
	// target := int(5)

	// result: 1
	// nums := []int{1, 3, 5, 6}
	// target := int(2)

	// result: 4
	// nums := []int{1, 3, 5, 6}
	// target := int(7)

	// result: 0
	// nums := []int{1, 3, 5, 6}
	// target := int(0)

	// result: 0
	// nums := []int{1}
	// target := int(0)

	// result: 0
	// nums := []int{-9, -8, -4, 0, 1, 3, 5, 6}
	// target := int(-9)

	// result: 6
	nums := []int{-9, -8, -4, 0, 1, 3, 5, 6}
	target := int(4)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := searchInsert(nums, target)
	fmt.Printf("result = %v\n", result)
}

