package main
import (
	"fmt"
)

func minPatches(nums []int, n int) int {
	miss := int(1)
	result := int(0)
	i := int(0)

	for miss <= n {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss += miss
			result++
		}
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{1,3}
	// n := int(6)

	// result: 2
	// nums := []int{1,5,10}
	// n := int(20)

	// result: 0
	nums := []int{1,2,2}
	n := int(5)

	// result: 
	// nums := []int{}
	// n := int(0)

	result := minPatches(nums, n)
	fmt.Printf("result = %v\n", result)
}

