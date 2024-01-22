package main
import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	n := len(nums)
	result := make([]int, 2)
	if n <= 1 {
		return result
	}

	m := make([]int, n)
	for _, num := range nums {
		m[num - 1]++ 
	}
	for i, count := range m {
		if count == 2 {
			result[0] = i + 1
		} else if count == 0 {
			result[1] = i + 1
		}
	}

	return result
}

func main() {
	// result: [2,3]
	// nums := []int{1,2,2,4}

	// result: [1,2]
	nums := []int{1,1}

	// result: 
	// nums := []int{}

	result := findErrorNums(nums)
	fmt.Printf("result = %v\n", result)
}

