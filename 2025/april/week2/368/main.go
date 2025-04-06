package main
import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    max := int(1)
    l := len(nums)
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1
    }
    for i := 1; i < l; i++ {
        for j := 0; j < i; j++ {
            if nums[i] % nums[j] == 0 && dp[j] + 1 > dp[i] {
                dp[i] = dp[j] + 1
                if dp[i] > max {
                    max = dp[i]
                }
            }
        }
    }

    result := []int{}
    prev := int(-1)
    for i := l - 1; i >= 0; i-- {
        if dp[i] == max && (prev == -1 || prev % nums[i] == 0) {
            result = append(result, nums[i])
            prev = nums[i]
            max--
        }
    }
    return result
}

func main() {
	// result: [1,2]
	// nums := []int{1,2,3}

	// result: [1,2,4,8]
	nums := []int{1,2,4,8}

	// result: 
	// nums := []int{}

	result := largestDivisibleSubset(nums)
	fmt.Printf("result = %v\n", result)
}
