package main
import (
	"fmt"
)

func findMatrix(nums []int) [][]int {
	freq := make([]int, len(nums) + 1)
	result := make([][]int, 0)
	for _, num := range nums {
		if freq[num] >= len(result) {
			result = append(result, make([]int, 0))
		}

		result[freq[num]] = append(result[freq[num]], num)
		freq[num]++
	}

	return result
}

func main() {
	// result: [[1,3,4,2],[1,3],[1]]
	// nums := []int{1,3,4,1,2,3,1}

	// result: [[4,3,2,1]]
	nums := []int{1,2,3,4}

	// result: 
	// nums := []int{}

	result := findMatrix(nums)
	fmt.Printf("result = %v\n", result)
}

