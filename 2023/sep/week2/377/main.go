package main
import (
	"fmt"
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

func main() {
	// result: 7
	// nums := []int{1,2,3}
	// target := int(4)

	// result: 0
	nums := []int{9}
	target := int(3)

	// result: 
	// nums := []int{}
	// target := int(0)

	result := combinationSum4(nums, target)
	fmt.Printf("result = %v\n", result)
}

