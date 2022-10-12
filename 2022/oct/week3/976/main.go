package main
import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		longest := nums[i]
		a := nums[i - 1]
		b := nums[i - 2]
		if longest >= a + b {
			continue
		}
		return longest + a + b
	}

	return 0
}

func main() {
	// result: 5
	// nums := []int{2,1,2}

	// result: 0
	// nums := []int{1,2,1}

	// result: 27
	// nums := []int{9,8,10,3,2,6,1,4,6,2,3}

	// result: 8
	nums := []int{3,6,2,3}

	// result: 
	// nums := []int{}

	result := largestPerimeter(nums)
	fmt.Printf("result = %v\n", result)
}

