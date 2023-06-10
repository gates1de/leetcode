package main
import (
	"fmt"
)

func maxValue(n int, index int, maxSum int) int {
	left := int(1)
	right := maxSum

	for left < right {
		mid := (left + right + 1) / 2
		if getSum(index, mid, n) <= maxSum {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func getSum(index int, value int, n int) int {
	count := int(0)
	if value > index {
		count += (value + value - index) * (index + 1) / 2
	} else {
		count += (value + 1) * value / 2 + index - value + 1
	}

	if value >= n - index {
		count += (value + value - n + 1 + index) * (n - index) / 2
	} else {
		count += (value + 1) * value / 2 + n - index - value
	}

	return count - value
}

func main() {
	// result: 2
	// n := int(4)
	// index := int(2)
	// maxSum := int(6)

	// result: 3
	n := int(6)
	index := int(1)
	maxSum := int(10)

	// result: 
	// n := int(0)
	// index := int(0)
	// maxSum := int(0)

	result := maxValue(n, index, maxSum)
	fmt.Printf("result = %v\n", result)
}

