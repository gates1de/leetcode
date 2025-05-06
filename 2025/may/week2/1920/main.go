package main
import (
	"fmt"
)

func buildArray(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = nums[num]
	}
	return result
}

func main() {
	// result: [0,1,2,4,5,3]
	// nums := []int{0,2,1,5,3,4}

	// result: [4,5,0,1,2,3]
	nums := []int{5,0,1,2,3,4}

	// result: []
	// nums := []int{}

	result := buildArray(nums)
	fmt.Printf("result = %v\n", result)
}

