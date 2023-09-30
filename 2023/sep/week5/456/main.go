package main
import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
    max := math.MinInt32
    stack := make([]int, 0)
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] < max {
            return true
        }
        for len(stack) > 0 && stack[len(stack) - 1] < nums[i] {
            max = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, nums[i])
    }
    return false
}

func main() {
	// result: false
	// nums := []int{1,2,3,4}

	// result: true
	// nums := []int{3,1,4,2}

	// result: true
	nums := []int{-1,3,2,0}

	// result: 
	// nums := []int{}

	result := find132pattern(nums)
	fmt.Printf("result = %v\n", result)
}

