package main
import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	m := map[int]int{}
	result := []int{}
	for _, num := range nums {
		if m[num] == 1 {
			result = append(result, num)
		}
		m[num]++
	}
	return result
}

func main() {
	// result: [2, 3]
	// nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	// result: [1]
	// nums := []int{1, 1, 2}

	// result: []
	nums := []int{1}

	// result: 
	// nums := []int{}

	result := findDuplicates(nums)
	fmt.Printf("result = %v\n", result)
}

