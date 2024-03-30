package main
import (
	"fmt"
)

func findDuplicate(nums []int) int {
	left := int(1)
	right := len(nums) - 1

	for left < right {
		mid := left + (right - left) / 2
		count := int(0)

		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	// result: 2
	// nums := []int{1,3,4,2,2}

	// result: 3
	// nums := []int{3,1,3,4,2}

	// result: 3
	nums := []int{3,3,3,3,3}

	// result: 
	// nums := []int{}

	result := findDuplicate(nums)
	fmt.Printf("result = %v\n", result)
}

