package main
import (
	"fmt"
)

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	if n < k {
		return make([]int, 0)
	} else if k == 1 {
		return nums
	}

	result := make([]int, n - k + 1)

	TOP:
	for i, _ := range result {
		maxNum := int(0)

		for j := i + 1; j < i + k; j++ {
			pre := nums[j - 1]
			num := nums[j]

			if pre + 1 != num {
				result[i] = -1
				continue TOP
			}

			maxNum = max(maxNum, num)
		}

		result[i] = maxNum
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
	// result: [3,4,-1,-1,-1]
	// nums := []int{1,2,3,4,3,2,5}
	// k := int(3)

	// result: [-1,-1]
	// nums := []int{2,2,2,2,2}
	// k := int(4)

	// result: [-1,3,-1,3,-1]
	// nums := []int{3,2,3,2,3,2}
	// k := int(2)

	// result: []
	nums := []int{1}
	k := int(1)

	result := resultsArray(nums, k)
	fmt.Printf("result = %v\n", result)
}

