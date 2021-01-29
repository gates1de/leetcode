package main
import (
	"fmt"
)

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}

	nums := make([]int, n + 1)
	for i, _ := range nums {
		nums[i] = -1
	}
	nums[0] = 0
	nums[1] = 1

	for i := 0; i < n + 1; i++ {
		even := 2 * i
		odd := 2 * i + 1
		if even < len(nums) && nums[even] < 0 && 2 <= even && even <= n {
			nums[even] = nums[i]
		}
		if odd < len(nums) && nums[odd] < 0 && 2 <= odd && odd <= n {
			nums[odd] = nums[i] + nums[i + 1]
		}
	}

	fmt.Printf("nums = %v\n", nums)
	return max(nums)
}

func max(target []int) int {
	result := int(0)
	for _, val := range target {
		if result < val {
			result = val
		}
	}
	return result
}

func main() {
	// n := int(0)
	// n := int(2)
	// n := int(3)
	n := int(7)
	result := getMaximumGenerated(n)
	fmt.Printf("result = %v\n", result)
}

