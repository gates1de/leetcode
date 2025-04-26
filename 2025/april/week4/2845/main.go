package main
import (
	"fmt"
)

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	cnt := make(map[int]int)
	result := int64(0)
	prefix := int(0)
	cnt[0] = 1

	for i := 0; i < n; i++ {
		if nums[i] % modulo == k {
			prefix++
		}

		result += int64(cnt[(prefix - k + modulo) % modulo])
		cnt[prefix % modulo]++
	}

	return result
}

func main() {
	// result: 3
	// nums := []int{3,2,4}
	// modulo := int(2)
	// k := int(1)

	// result: 2
	nums := []int{3,1,9,6}
	modulo := int(3)
	k := int(0)

	// result: 
	// nums := []int{}
	// modulo := int(0)
	// k := int(0)

	result := countInterestingSubarrays(nums, modulo, k)
	fmt.Printf("result = %v\n", result)
}

