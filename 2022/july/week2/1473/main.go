package main
import (
	"fmt"
	"math"
)

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, 100)
	for i, _ := range dp {
		dp[i] = make([][]int, 100)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, 21)
			for k, _ := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	result := helper(houses, cost, target, 0, 0, 0, dp)
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func helper(houses []int, cost [][]int, target int, index int, count int, prevHouseColor int, dp [][][]int) int {
	if index == len(houses) {
		// If all houses are traversed, check if the neighbor count is as expected or not
		if count == target {
			return 0
		}
		return math.MaxInt32
	}

	// If the neighborhoods are more than the threshold, we can't have target neighborhoods
	if count > target {
		return math.MaxInt32
	}

	// We have already calculated the answer so no need to go into recursion
	if dp[index][count][prevHouseColor] != -1 {
		return dp[index][count][prevHouseColor]
	}
        
	minCost := math.MaxInt32
	// If the house is already painted, update the values accordingly
	if houses[index] != 0 {
		newCount := count
		if houses[index] != prevHouseColor {
			newCount++
		}
		minCost = helper(houses, cost, target, index + 1, newCount, houses[index], dp)
	} else {
		totalColors := len(cost[0])
		// If the house is not painted, try every possible color and store the minimum cost
		for color := 1; color <= totalColors; color++ {
			newCount := count
			if color != prevHouseColor {
				newCount++
			}
			currentCost := cost[index][color - 1] + helper(houses, cost, target, index + 1, newCount, color, dp)

			if currentCost < minCost {
				minCost = currentCost
			}
		}
	}
		
	// Return the minimum cost and also storing it for future reference (memoization)
	dp[index][count][prevHouseColor] = minCost
	return minCost
}

func main() {
	// result: 9
	// houses := []int{0,0,0,0,0}
	// cost := [][]int{{1,10},{10,1},{10,1},{1,10},{5,1}}
	// m := int(5)
	// n := int(2)
	// target := int(3)

	// result: 11
	// houses := []int{0,2,1,2,0}
	// cost := [][]int{{1,10},{10,1},{10,1},{1,10},{5,1}}
	// m := int(5)
	// n := int(2)
	// target := int(3)

	// result: -1
	houses := []int{3,1,2,3}
	cost := [][]int{{1,1,1},{1,1,1},{1,1,1},{1,1,1}}
	m := int(4)
	n := int(3)
	target := int(3)

	// result: 
	// houses := []int{}
	// cost := [][]int{}
	// m := int(0)
	// n := int(0)
	// target := int(0)

	result := minCost(houses, cost, m, n, target)
	fmt.Printf("result = %v\n", result)
}

