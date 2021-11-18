package main
import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	m := map[int]bool{}

	for _, n := range nums {
		m[n] = true
	}
	for i := 1; i <= len(nums); i++ {
		if !m[i] {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// result: [5,6]
	// nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	// result: [2]
	// nums := []int{1, 1}

	// result: [5,10,11,12,13,14,15,16,17,18,19,20,21]
	nums := []int{3,4,9,2,6,1,7,8,3,6,4,9,8,1,2,3,6,4,9,2,8}

	// result: 
	// nums := []int{}

	result := findDisappearedNumbers(nums)
	fmt.Printf("result = %v\n", result)
}

