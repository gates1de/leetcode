package main
import (
	"fmt"
	"sort"
)

func reductionOperations(nums []int) int {
	sort.Ints(nums)
	result := int(0)
	up := int(0)

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i - 1] {
			up++
		}

		result += up
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{5,1,3}

	// result: 0
	// nums := []int{1,1,1}

	// result: 4
	nums := []int{1,1,2,2,3}

	// result: 
	// nums := []int{}

	result := reductionOperations(nums)
	fmt.Printf("result = %v\n", result)
}

