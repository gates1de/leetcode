package main

import (
	"fmt"
)

func separateDigits(nums []int) []int {
	n := len(nums)
	result := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		for num > 0 {
			result = append(result, num % 10)
			num /= 10
		}
	}

	for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func main() {
  // result: [7,1,3,9]
	// nums := []int{7,1,3,9}

	// result: [1,3,2,5,8,3,7,7]
	nums := []int{13,25,83,77}

	// result: 
	// nums := []int{}

	result := separateDigits(nums)
	fmt.Printf("result = %v\n", result)
}

