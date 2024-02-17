package main
import (
	"fmt"
)

func majorityElement(nums []int) int {
	m := make(map[int]int)
	maxCount := int(0)
	result := int(-1)

	for _, num := range nums {
		m[num]++
		if maxCount < m[num] {
			maxCount = m[num]
			result = num
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// nums := []int{3,2,3}

	// result: 2
	nums := []int{2,2,1,1,1,2,2}

	// result: 
	// nums := []int{}

	result := majorityElement(nums)
	fmt.Printf("result = %v\n", result)
}

