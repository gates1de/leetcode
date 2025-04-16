package main
import (
	"fmt"
)

func countGood(nums []int, k int) int64 {
	n := len(nums)
	same := int(0)
	j := int(-1)
	counter := make(map[int]int)
	result := int64(0)

	for i := 0; i < n; i++ {
		for same < k && j + 1 < n {
			j++
			same += counter[nums[j]]
			counter[nums[j]]++
		}

		if same >= k {
			result += int64(n - j)
		}

		counter[nums[i]]--
		same -= counter[nums[i]]
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{1,1,1,1,1}
	// k := int(10)

	// result: 4
	nums := []int{3,1,4,3,2,2,4}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := countGood(nums, k)
	fmt.Printf("result = %v\n", result)
}

