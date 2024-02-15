package main
import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	preSum := int64(0)
	result := int64(-1)
	for _, num := range nums {
		num64 := int64(num)
		if num64 < preSum {
			result = num64 + preSum
		}
		preSum += num64
	}

	return result
}

func main() {
	// result: 15
	// nums := []int{5,5,5}

	// result: 12
	// nums := []int{1,12,1,2,5,50,3}

	// result: -1
	nums := []int{5,5,50}

	// result: 
	// nums := []int{}

	result := largestPerimeter(nums)
	fmt.Printf("result = %v\n", result)
}

