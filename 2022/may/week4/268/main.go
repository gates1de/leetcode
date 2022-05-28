package main
import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}

	return len(nums)
}

func main() {
	// result: 2
	// nums := []int{3, 0, 1}

	// result: 2
	// nums := []int{0, 1}

	// result: 8
	nums := []int{9,6,4,2,3,5,7,0,1}

	// result: 
	// nums := []int{}

	result := missingNumber(nums)
	fmt.Printf("result = %v\n", result)
}

