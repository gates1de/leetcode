package main
import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)
	left := int(0)
	right := nums[n - 1] - nums[0]

	for left < right {
		mid := left + (right - left) / 2

		if countValidPairs(nums, mid) >= p {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func countValidPairs(nums []int, threshold int) int {
	index := int(0)
	count := int(0)
	for index < len(nums) - 1 {
		if nums[index + 1] - nums[index] <= threshold {
			count++
			index++
		}
		index++
	}
	return count
}

func main() {
	// result: 1
	// nums := []int{10,1,2,7,1,3}
	// p := int(2)

	// result: 0
	nums := []int{4,2,1,2}
	p := int(1)

	// result: 
	// nums := []int{}
	// p := int(0)

	result := minimizeMax(nums, p)
	fmt.Printf("result = %v\n", result)
}

