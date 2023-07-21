package main
import (
	"fmt"
)

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	length := make([]int, n)
	counts := make([]int, n)
	for i, _ := range length {
		length[i] = 1
		counts[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if length[j] + 1 > length[i] {
					length[i] = length[j] + 1
					counts[i] = 0
				}

				if length[j] + 1 == length[i] {
					counts[i] += counts[j]
				}
			}
		}
	}

	maxLength := int(0)
	result := int(0)
	
	for _, l := range length {
		maxLength = max(maxLength, l)
	}

	for i, l := range length {
		if l == maxLength {
			result += counts[i]
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
	// result: 2
	// nums := []int{1,3,5,4,7}

	// result: 5
	nums := []int{2,2,2,2,2}

	// result: 
	// nums := []int{}

	result := findNumberOfLIS(nums)
	fmt.Printf("result = %v\n", result)
}

