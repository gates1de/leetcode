package main
import (
	"fmt"
)

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	backtrack(make([]int, 0), nums, &result)
	return result
}

func backtrack(current []int, nums []int, result *[][]int) {
	if len(current) == len(nums) {
		*result = append(*result, current)
		return
	}

	for _, num := range nums {
		if contains(num, current) {
			continue
		}

		newCurrent := make([]int, len(current) + 1)
		copy(newCurrent, current)
		newCurrent[len(current)] = num
		backtrack(newCurrent, nums, result)
	}
}

func contains(target int, nums []int) bool {
	for _, num := range nums {
		if target == num {
			return true
		}
	}
	return false
}

func main() {
	// result: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
	// nums := []int{1,2,3}

	// result: [[0,1],[1,0]]
	// nums := []int{0,1}

	// result: [[1]]
	nums := []int{1}

	// result: 
	// nums := []int{}

	result := permute(nums)
	fmt.Printf("result = %v\n", result)
}

