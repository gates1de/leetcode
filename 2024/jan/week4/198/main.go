package main
import (
	"fmt"
)

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	s0 := make([]int, 2)
	s1 := make([]int, 2)
	s1[0] = nums[0]

	for i := 1; i < n; i++ {
		s0[i % 2] = max(s0[(i - 1) % 2], s1[(i - 1) % 2])
		s1[i % 2] = s0[(i - 1) % 2] + nums[i]
	}

	return max(s0[(n - 1) % 2], s1[(n - 1) % 2])
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// nums := []int{1,2,3,1}

	// result: 12
	nums := []int{2,7,9,3,1}

	// result: 
	// nums := []int{}

	result := rob(nums)
	fmt.Printf("result = %v\n", result)
}

