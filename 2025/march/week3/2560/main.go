package main
import (
	"fmt"
)

func minCapability(nums []int, k int) int {
	n := len(nums)
	maxReward := int(0)
	result := int(1)

	for _, num := range nums {
		maxReward = max(maxReward, num)
	}

	for result < maxReward {
		mid := (result + maxReward) / 2
		possibleThefts := int(0)

		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				possibleThefts++
				i++
			}
		}

		if possibleThefts >= k {
			maxReward = mid
		} else {
			result = mid + 1
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// nums := []int{2,3,5,9}
	// k := int(2)

	// result: 2
	nums := []int{2,7,9,3,1}
	k := int(2)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := minCapability(nums, k)
	fmt.Printf("result = %v\n", result)
}

