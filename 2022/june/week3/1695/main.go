package main
import (
	"fmt"
)

func maximumUniqueSubarray(nums []int) int {
	m := map[int]int{}
	result := int(0)

	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = nums[i] + sums[i - 1]
	}

	left := int(0)
	for i, num := range nums {
		if removeIndex, ok := m[num]; ok {
			left = max(left, removeIndex + 1)
			delete(m, num)
		}

		m[num] = i
		if left == 0 {
			result = max(result, sums[i])
			continue
		}

		currentSum := sums[i] - sums[left - 1]
		result = max(result, currentSum)
	}

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	// result: 17
	// nums := []int{4, 2, 4, 5, 6}

	// result: 8
	nums := []int{5, 2, 1, 2, 5, 2, 1, 2, 5}

	// result: 
	// nums := []int{}

	result := maximumUniqueSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

