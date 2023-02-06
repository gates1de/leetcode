package main
import (
	"fmt"
)

func shuffle(nums []int, n int) []int {
	result := make([]int, 2 * n)
	for i := 0; i < n; i++ {
		result[i * 2] = nums[i]
		result[1 + (i * 2)] = nums[n + i]
	}
	return result
}

func main() {
	// result: [2,3,5,4,1,7]
	// nums := []int{2,5,1,3,4,7}
	// n := int(3)

	// result: [1,4,2,3,3,2,4,1]
	// nums := []int{1,2,3,4,4,3,2,1}
	// n := int(4)

	// result: [1,2,1,2]
	nums := []int{1,1,2,2}
	n := int(2)

	// result: 
	// nums := []int{}
	// n := int(0)

	result := shuffle(nums, n)
	fmt.Printf("result = %v\n", result)
}

