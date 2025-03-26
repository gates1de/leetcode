package main
import (
	"fmt"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	m := len(grid)
	n := len(grid[0])
	nums := make([]int, 0, m * n)
	result := int(0)

	for row, _ := range grid {
		for col, _ := range grid[row] {
			nums = append(nums, grid[row][col])
		}
	}

	sort.Ints(nums)
	finalCommonNumber := nums[len(nums) / 2]

	for _, num := range nums {
		if num % x != finalCommonNumber % x {
			return -1
		}

		result += abs(finalCommonNumber - num) / x
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 4
	// grid := [][]int{{2,4},{6,8}}
	// x := int(2)

	// result: 5
	// grid := [][]int{{1,5},{2,3}}
	// x := int(1)

	// result: -1
	grid := [][]int{{1,2},{3,4}}
	x := int(2)

	// result: 
	// grid := [][]int{}
	// x := int(0)

	result := minOperations(grid, x)
	fmt.Printf("result = %v\n", result)
}

