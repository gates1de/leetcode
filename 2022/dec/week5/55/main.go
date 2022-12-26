package main
import (
	"fmt"
)

func canJump(nums []int) bool {
    memo := make([]int, len(nums))
    memo[len(nums) - 1] = 1
    for i := len(nums) - 2; i >= 0; i-- {
        furthestJump := min(i + nums[i], len(nums) - 1)
        for j := i + 1; j <= furthestJump; j++ {
            if memo[j] == 1 {
                memo[i] = 1
                break
            }
        }
    }
    return memo[0] == 1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	// result: true
	// nums := []int{2,3,1,1,4}

	// result: false
	nums := []int{3,2,1,0,4}

	// result: 
	// nums := []int{}

	result := canJump(nums)
	fmt.Printf("result = %v\n", result)
}

