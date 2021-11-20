package main
import (
	"fmt"
)

func singleNonDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return nums[0]
	}

	mid := (len(nums) - 1) / 2
	if nums[mid] != nums[mid + 1] {
		mid++
	}

	if len(nums[:mid]) % 2 == 0 {
		return singleNonDuplicate(nums[mid:])
	}
	return singleNonDuplicate(nums[:mid])
}

func main() {
	// result: 2
	// nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}

	// result: 10
	// nums := []int{3, 3, 7, 7, 10, 11, 11}

	// result: 4
	// nums := []int{1, 1, 2, 2, 3, 3, 4, 8, 8}

	// result: 0
	// nums := []int{0}

	// result: 1
	nums := []int{0, 0, 1}

	// result: 
	// nums := []int{}

	result := singleNonDuplicate(nums)
	fmt.Printf("result = %v\n", result)
}

