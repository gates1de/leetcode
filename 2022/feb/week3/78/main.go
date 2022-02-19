package main
import (
	"fmt"
)

func subsets(nums []int) [][]int {
	result := [][]int{[]int{}}
	for i, num := range nums {
		helper([]int{num}, nums[i + 1:], &result)
	}
	return result
}

func helper(current []int, nums []int, result *[][]int) {
	if len(nums) == 0 {
		// fmt.Println(current)
		*result = append(*result, current)
		return
	}

	num := nums[0]
	helper(current, nums[1:], result)
	newCurrent := make([]int, len(current) + 1)
	copy(newCurrent, current)
	newCurrent[len(current)] = num
	helper(newCurrent, nums[1:], result)
}

func main() {
	// result: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
	// nums := []int{1, 2, 3}

	// result: [[],[0]]
	// nums := []int{0}

	// result: 
	nums := []int{-10, -9, 0, 9, 10}

	// result: 
	// nums := []int{}

	result := subsets(nums)
	fmt.Printf("result = %v\n", result)
}

