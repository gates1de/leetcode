package main
import (
	"fmt"
)

func minimumIndex(nums []int) int {
	n := len(nums)
	x := nums[0]
	count := int(0)
	xCount := int(0)

	for _, num := range nums {
		if num == x {
			count++
		} else {
			count--
		}

		if count == 0 {
			x = num
			count = 1
		}
	}

	for _, num := range nums {
		if num == x {
			xCount++
		}
	}
	
	count = 0
	for i, num := range nums {
		if num == x {
			count++
		}

		remainingCount := xCount - count
		if count * 2 > i + 1 && remainingCount * 2 > n - i - 1 {
			return i
		}
	}

	return -1
}

func main() {
	// result: 2
	// nums := []int{1,2,2,2}

	// result: 4
	// nums := []int{2,1,3,1,1,1,7,1,2,1}

	// result: -1
	nums := []int{3,3,3,3,7,2,2}

	// result: 
	// nums := []int{}

	result := minimumIndex(nums)
	fmt.Printf("result = %v\n", result)
}

