package main
import (
	"fmt"
)

func countCompleteSubarrays(nums []int) int {
	n := len(nums)
	counter := make(map[int]int)
	distinct := make(map[int]bool)
	for _, num := range nums {
		distinct[num] = true
	}
	distinctCount := len(distinct)

	right := int(0)
	result := int(0)
	for left := 0; left < n; left++ {
		if left > 0 {
			removeNum := nums[left - 1]
			counter[removeNum]--
			if counter[removeNum] == 0 {
				delete(counter, removeNum)
			}
		}

		for right < n && len(counter) < distinctCount {
			addNum := nums[right]
			counter[addNum]++
			right++
		}

		if len(counter) == distinctCount {
			result += (n - right + 1)
		}
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{1,3,1,2,2}

	// result: 10
	nums := []int{5,5,5,5}

	// result: 
	// nums := []int{}

	result := countCompleteSubarrays(nums)
	fmt.Printf("result = %v\n", result)
}

