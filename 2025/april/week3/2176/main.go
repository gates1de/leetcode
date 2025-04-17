package main
import (
	"fmt"
)

func countPairs(nums []int, k int) int {
	n := len(nums)
	result := int(0)

	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			if (i * j) % k == 0 && nums[i] == nums[j] {
				result++
			}
		}
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{3,1,2,2,2,1,3}
	// k := int(2)

	// result: 0
	nums := []int{1,2,3,4}
	k := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := countPairs(nums, k)
	fmt.Printf("result = %v\n", result)
}

