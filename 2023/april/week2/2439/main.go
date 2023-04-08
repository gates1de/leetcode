package main
import (
	"fmt"
)

func minimizeArrayValue(nums []int) int {
	result := float64(0)
	prefixSum := float64(0)

	for i, num := range nums {
		prefixSum += float64(num)
		result = max(result, (prefixSum + float64(i)) / float64(i + 1))
	}

	return int(result)
}

func max(a, b float64) float64 {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// nums := []int{3,7,1,6}

	// result: 10
	nums := []int{10,1}

	// result: 
	// nums := []int{}

	result := minimizeArrayValue(nums)
	fmt.Printf("result = %v\n", result)
}

