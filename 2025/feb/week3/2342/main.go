package main
import (
	"fmt"
)

func maximumSum(nums []int) int {
	result := int(-1)
	m := make([]int, 82)

	for _, num := range nums {
		digitSum := int(0)

		for val := num; val != 0; val /= 10 {
			digit := val % 10
			digitSum += digit
		}

		if m[digitSum] > 0 {
			result = max(result, m[digitSum] + num)
		}
		m[digitSum] = max(m[digitSum], num)
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
	// result: 54
	// nums := []int{18,43,36,13,7}

	// result: -1
	nums := []int{10,12,19,14}

	// result: 
	// nums := []int{}

	result := maximumSum(nums)
	fmt.Printf("result = %v\n", result)
}

