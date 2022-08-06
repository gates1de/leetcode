package main
import (
	"fmt"
	"sort"
)

func combinationSum4(nums []int, target int) int {
    dp := make([]int, target + 1)
    dp[0] = 1
    for i := 1; i <= target; i++ {
        for _, num := range nums {
            if i - num >= 0 {
                dp[i] += dp[i - num]
            }
        }
    }
    return dp[target]
}

// Time Limit Exceeded
func ngSolution(nums []int, target int) int {
	sort.Ints(nums)
	m := map[int]int{}
	for i := 1; i <= target; i++ {
		if i < nums[0] {
			continue
		}
		helper(i, 0, nums, i, m)
	}
	return m[target]
}

func helper(remaining int, count int, nums []int, target int, m map[int]int) {
	if remaining == 0 {
		m[target] += count
		return
	}

	if val, ok := m[remaining]; ok {
		m[target] += val * count
		return
	}

	if remaining < 0 {
		return
	}

	for _, num := range nums {
		if remaining - num < 0 {
			break
		}
		helper(remaining - num, count + 1, nums, target, m)
	}
}

func main() {
	// result: 7
	// nums := []int{1, 2, 3}
	// target := int(4)

	// result: 0
	// nums := []int{9}
	// target := int(3)

	// result: 39882198
	nums := []int{4,2,1}
	target := int(32)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := combinationSum4(nums, target)
	fmt.Printf("result = %v\n", result)
}

