package main
import (
	"fmt"
)

func countSmaller(nums []int) []int {
	result := make([]int, len(nums))
	for i, n1 := range nums {
		for _, n2 := range nums[i + 1:] {
			if n1 > n2 {
				result[i]++
			}
		}
	}
	return result
}

func main() {
	// result: [2,1,1,0]
	// nums := []int{5, 2, 6, 1}

	// result: [0]
	// nums := []int{-1}

	// result: [0, 0]
	nums := []int{-1, -1}

	// result: 
	// nums := []int{}

	result := countSmaller(nums)
	fmt.Printf("result = %v\n", result)
}

