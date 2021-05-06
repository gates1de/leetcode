package main
import (
	"fmt"
)

func jump(nums []int) int {
	result := int(0)
	curEnd := int(0)
	curFarthest := int(0)

	for i := 0; i < len(nums) - 1; i++ {
		if curFarthest < i + nums[i] {
			curFarthest = i + nums[i]
		}
		if i == curEnd {
			result++
			curEnd = curFarthest
		}
	}
	return result
}

// Wrong Answer
func ngSolution(nums []int) int {
	goal := len(nums) - 1
	result := int(0)
	i := int(0)
	for i < goal {
		dist := nums[i]
		if i + 1 + dist > goal {
			result++
			return result
		}
		maxJumpIndex := min(i + 1 + dist, goal + 1)
		nextIndex := i + maxIndex(nums[i + 1:maxJumpIndex], dist) + 1
		if nextIndex == -1 {
			i += dist
		} else {
			i = nextIndex
		}
		fmt.Printf("dist = %v, nextIndex = %v\n", dist, nextIndex)
		result++
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func maxIndex(nums []int, dist int) int {
	fmt.Printf("nums = %v\n", nums)
	maxNum := int(0)
	result := int(-1)
	for i, num := range nums {
		if maxNum <= num && dist < i + 1 + num {
			result = i
			maxNum = num
		}
	}
	return result
}

func main() {
	// result: 2
	// nums := []int{2, 3, 1, 1, 4}

	// result: 2
	// nums := []int{2, 3, 0, 1, 4}

	// result: 1
	// nums := []int{3, 2, 1}

	// result: 2
	// nums := []int{1, 2, 3}

	// result: 3
	// nums := []int{1, 2, 1, 1, 1}

	// result: 2
	// nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}

	// result: 2
	nums := []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}

	// result: 
	// nums := []int{}

	result := jump(nums)
	fmt.Printf("result = %v\n", result)
}

