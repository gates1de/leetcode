package main
import (
	"fmt"
)

func maximumCount(nums []int) int {
	positiveCount := int(0)
	negativeCount := int(0)

	for i, num := range nums {
		if num == 0 {
			continue
		} else if num > 0 {
			positiveCount = len(nums[i:])
			break
		} else {
			negativeCount = len(nums[:i + 1])
		}
	}

	return max(positiveCount, negativeCount)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// nums := []int{-2,-1,-1,1,2,3}

	// result: 3
	// nums := []int{-3,-2,-1,0,0,1,2}

	// result: 4
	nums := []int{5,20,66,1314}

	// result: 
	// nums := []int{}

	result := maximumCount(nums)
	fmt.Printf("result = %v\n", result)
}

