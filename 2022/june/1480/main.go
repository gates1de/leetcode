package main
import (
	"fmt"
)

func runningSum(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i - 1]
    }
    return nums
}

func main() {
	// result: [1,3,6,10]
	nums := []int{1, 2, 3, 4}

	// result: [1]
	// nums := []int{1}

	// result: 
	// nums := []int{}

	result := runningSum(nums)
	fmt.Printf("result = %v\n", result)
}

