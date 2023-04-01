package main
import (
	"fmt"
)

func search(nums []int, target int) int {
    return binarySearch(0, nums, target)
}

func binarySearch(index int, nums []int, target int) int {
    if index >= len(nums) {
        return -1
    }

    if nums[index] == target {
        return index
    }

    left := binarySearch(2 * index + 1, nums, target)
    if left >= 0 {
        return left
    }

    return binarySearch(2 * index + 2, nums, target)
}

func main() {
	// result: 4
	// nums := []int{-1,0,3,5,9,12}
	// target := int(9)

	// result: -1
	nums := []int{-1,0,3,5,9,12}
	target := int(2)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := search(nums, target)
	fmt.Printf("result = %v\n", result)
}

