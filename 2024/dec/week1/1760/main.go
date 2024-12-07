package main
import (
	"fmt"
	"math"
)

func minimumSize(nums []int, maxOperations int) int {
	left := int(1)
	right := int(0)

	for _, num := range nums {
		right = max(right, num)
	}

	for left < right {
		mid := (left + right) / 2

		if isValid(mid, nums, maxOperations) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isValid(maxBallsInBag int, nums []int, maxOperations int) bool {
	totalOperations := int(0)

	for _, num := range nums {
		operations := math.Ceil(float64(num) / float64(maxBallsInBag)) - 1
		totalOperations += int(operations)

		if totalOperations > maxOperations {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// nums := []int{9}
	// maxOperations := int(2)

	// result: 2
	nums := []int{2,4,8,2}
	maxOperations := int(4)

	// result: 
	// nums := []int{}
	// maxOperations := int(0)

	result := minimumSize(nums, maxOperations)
	fmt.Printf("result = %v\n", result)
}

