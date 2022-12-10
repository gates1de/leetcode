package main
import (
	"fmt"
	"math"
)

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sum := int(0)
	for _, num := range nums {
		sum += num
	}

	if sum == 0 {
		return 0
	}

	minDiff := math.MaxInt32
	leftSum := int(0)
	rightSum := sum
	result := int(0)
	for i := 0; i < n; i++ {
		num := nums[i]
		leftSum += num
		rightSum -= num
		numOfLeftNums := i + 1
		numOfRightNums := n - i - 1

		if numOfRightNums == 0 {
			break
		}
		diff := abs(leftSum / numOfLeftNums, rightSum / numOfRightNums)

		if diff < minDiff {
			minDiff = diff
			result = i
		}
	}

	if abs(sum / n, 0) < minDiff {
		result = n - 1
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: 3
	// nums := []int{2,5,3,9,5,3}

	// result: 0
	// nums := []int{0}

	// result: 0
	// nums := []int{9,9,9,9,9,9,9,9,8}

	// result: 0
	// nums := []int{0,0,0,0,0}

	// result: 0
	nums := []int{0,1,0,1,0,1}

	// result: 
	// nums := []int{}

	result := minimumAverageDifference(nums)
	fmt.Printf("result = %v\n", result)
}

