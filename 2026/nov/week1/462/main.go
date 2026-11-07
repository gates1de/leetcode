package main

import (
	"fmt"
)

func minMoves2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	median := quickSelect(nums, len(nums)/2)
	result := 0
	for _, num := range nums {
		if num < median {
			result += median - num
		} else {
			result += num - median
		}
	}

	return result
}

func quickSelect(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] < nums[left] {
			nums[middle], nums[left] = nums[left], nums[middle]
		}
		if nums[right] < nums[left] {
			nums[right], nums[left] = nums[left], nums[right]
		}
		if nums[right] < nums[middle] {
			nums[right], nums[middle] = nums[middle], nums[right]
		}

		pivot := nums[middle]
		i := left
		j := right
		for i <= j {
			for nums[i] < pivot {
				i++
			}
			for nums[j] > pivot {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}

		if target <= j {
			right = j
		} else if target >= i {
			left = i
		} else {
			return nums[target]
		}
	}

	return nums[target]
}

func main() {
	// result: 2
	// nums := []int{1, 2, 3}

	// result: 16
	// nums := []int{1, 10, 2, 9}

	// result:
	nums := []int{3, -20, 8, 14, -6, 2, 30, 4, 2, 6, 10, -4, -6, 7, 2, -30, 1, 8, 46, 8, 3, 0}

	// result:
	// nums := []int{}

	result := minMoves2(nums)
	fmt.Printf("result = %v\n", result)
}
