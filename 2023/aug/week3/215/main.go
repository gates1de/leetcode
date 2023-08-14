package main
import (
	"fmt"
	"math"
)

func findKthLargest(nums []int, k int) int {
	minValue := math.MaxInt32
	maxValue := math.MinInt32
	for _, num := range nums {
		minValue = min(minValue, num)
		maxValue = max(maxValue, num)
	}

	count := make([]int, maxValue - minValue + 1)
	for _, num := range nums {
		count[num - minValue]++
	}

	remain := k
	for num := len(count) - 1; num >= 0; num-- {
		remain -= count[num]
		if remain <= 0 {
			return num + minValue
		}
	}

	return -1
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
	// result: 5
	// nums := []int{3,2,1,5,6,4}
	// k := int(2)

	// result: 4
	nums := []int{3,2,3,1,2,4,5,5,6}
	k := int(4)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := findKthLargest(nums, k)
	fmt.Printf("result = %v\n", result)
}

