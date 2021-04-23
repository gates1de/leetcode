package main
import (
	"fmt"
	"math"
)

func leastBricks(wall [][]int) int {
	dp := map[int]int{}
	for _, nums := range wall {
		sum := int(0)
		for _, num := range nums {
			sum += num
			dp[sum]++
		}
	}

	wallWidth := sum(wall[0])
	dp[wallWidth] = 0

	maxWallThroughCount := int(0)
	for _, num := range dp {
		if maxWallThroughCount < num {
			maxWallThroughCount = num
		}
	}
	// fmt.Printf("dp = %v\n", dp)
	return len(wall) - maxWallThroughCount
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: 2
	// wall := [][]int{
	// 	{1, 2, 2, 1},
	// 	{3, 1, 2},
	// 	{1, 3, 2},
	// 	{2, 4},
	// 	{3, 1, 2},
	// 	{1, 3, 1, 1},
	// }

	// result: 3
	// wall := [][]int{
	// 	{1},
	// 	{1},
	// 	{1},
	// }

	// result: 3
	// wall := [][]int{
	// 	{100000000},
	// 	{100000000},
	// 	{100000000},
	// }

	// result: 0
	// wall := [][]int{
	// 	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	// 	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	// 	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	// }

	// result: 0
	// wall := [][]int{
	// 	{7, 1, 1, 1},
	// 	{1, 4, 2, 3},
	// 	{7, 3},
	// }

	// result: 17
	wall := [][]int{
		{6}, {6}, {2, 4}, {6}, {1, 2, 2, 1}, {6}, {2, 1, 2, 1}, {1, 5}, {4, 1, 1}, {1, 4, 1}, {4, 2}, {3, 3}, {2, 2, 2}, {5, 1}, {5, 1}, {6}, {4, 2}, {1, 5},
		{2, 3, 1}, {4, 2}, {1, 1, 4}, {1, 3, 1, 1}, {2, 3, 1}, {3, 3}, {3, 1, 1, 1}, {3, 2, 1}, {6}, {3, 2, 1}, {1, 5}, {1, 4, 1},
	}

	// result: 
	// wall := [][]int{
	// 	{},
	// 	{},
	// }

	result := leastBricks(wall)
	fmt.Printf("result = %v\n", result)
}

