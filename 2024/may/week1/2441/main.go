package main
import (
	"fmt"
)

func findMaxK(nums []int) int {
	result := int(-1)
	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		absNum := abs(num)
		if m[-num] && result < absNum {
			result = absNum
		}

		m[num] = true
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 3
	// nums := []int{-1,2,-3,3}

	// result: 7
	// nums := []int{-1,10,6,7,-7,1}

	// result: -1
	nums := []int{-10,8,6,7,-2,-3}

	// result: 
	// nums := []int{}

	result := findMaxK(nums)
	fmt.Printf("result = %v\n", result)
}

