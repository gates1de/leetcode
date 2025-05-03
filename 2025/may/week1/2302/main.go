package main
import (
	"fmt"
)

func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	result := int64(0)
	total := int64(0)

	for i, j := 0, 0; j < n; j++ {
		total += int64(nums[j])

		for i <= j && total * int64(j - i + 1) >= k {
			total -= int64(nums[i])
			i++
		}

		result += int64(j - i + 1)
	}

	return result
}

func main() {
	// result: 6
	// nums := []int{2,1,4,3,5}
	// k := int64(10)

	// result: 5
	nums := []int{1,1,1}
	k := int64(5)

	// result: 
	// nums := []int{}
	// k := int64(0)

	result := countSubarrays(nums, k)
	fmt.Printf("result = %v\n", result)
}

