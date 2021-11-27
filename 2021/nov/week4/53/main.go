package main
import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
    prevMax := 0
    globalMax := -(1 << 31)
    for _, val := range nums {
        curMax := val + prevMax
        if curMax > globalMax {
            globalMax = curMax
        }
        prevMax = curMax
        if prevMax < 0 {
            prevMax = 0
        }
    }

    return globalMax
}

// Slow & Use more memory
func mySolution(nums []int) int {
	current := int(0)
	result := math.MinInt32
	for _, num := range nums {
		if current < 0 && current <= num {
			current = num
			if result < current {
				result = current
			}
			continue
		}

		current += num
		if result < current {
			result = current
		}
	}
	return result
}

func main() {
	// result: 6
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	// result: 1
	// nums := []int{1}

	// result: 23
	// nums := []int{5, 4, -1, 7, 8}

	// result: 39
	// nums := []int{-3,2,9,-8,1,4,6,3,-7,-8,4,-1,6,-9,3,8,2,6,1,8,9,-6,-4,3}

	// result: -1
	nums := []int{-1}

	// result: 
	// nums := []int{}

	result := maxSubArray(nums)
	fmt.Printf("result = %v\n", result)
}

