package main
import (
	"fmt"
	"math"
	"sort"
)


func minOperations(nums []int, x int) int {
    result := int(-1)
    numsSum := sum(nums)
    maxSubarraySum := numsSum - x

    currentMax := int(0)
    j := int(0)
    for i, val := range nums {
        currentMax += val
        for currentMax > maxSubarraySum && i >= j {
            currentMax -= nums[j]
            j++
        }

        if maxSubarraySum == currentMax {
            result = max(result, i - j + 1)
        }
    }

    if result == -1 {
        return result
    }
    return len(nums) - result
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func sum(nums []int) int {
    result := int(0)
    for _, val := range nums {
        result += val
    }
    return result
}

func main() {
	// result: 2
	// nums := []int{1, 1, 4, 2, 3}
	// x := int(5)

	// result: -1
	// nums := []int{5, 6, 7, 8, 9}
	// x := int(4)

	// result: 5
	// nums := []int{3, 2, 20, 1, 1, 3}
	// x := int(10)

	// result: 1
	// nums := []int{7, 1, 1, 1, 4, 3}
	// x := int(7)

	// result: 2
	nums := []int{5, 1, 1, 1, 4, 3}
	x := int(7)

	// result: 
	// nums := []int{}
	// x := int(0)

	result := minOperations(nums, x)
	fmt.Printf("result = %v\n", result)
}

