package main
import (
	"fmt"
)


func countSubarrays(nums []int, minK int, maxK int) int64 {
    n := len(nums)
    result := int64(0)
    minPosition := int(-1)
    maxPosition := int(-1)
    leftBound := int(-1)

    for i := 0; i < n; i++ {
        num := nums[i]
        if num < minK || num > maxK {
            leftBound = i
        }

        if num == minK {
            minPosition = i
        }
        if num == maxK {
            maxPosition = i
        }

        result += int64(max(0, min(maxPosition, minPosition) - leftBound))
    }

    return result
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func main() {
	// result: 2
	// nums := []int{1,3,5,2,7,5}
	// minK := int(1)
	// maxK := int(5)

	// result: 10
	nums := []int{1,1,1,1}
	minK := int(1)
	maxK := int(1)

	// result: 
	// nums := []int{}
	// minK := int(0)
	// maxK := int(0)

	result := countSubarrays(nums, minK, maxK)
	fmt.Printf("result = %v\n", result)
}

