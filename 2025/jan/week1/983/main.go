package main
import (
	"fmt"
	"math"
)

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, len(days))
	for i, _ := range dp {
		dp[i] = -1
	}
	return helper(0, days, costs, dp)
}

func helper(day int, days []int, costs []int, dp []int) int {
	if day >= len(days) {
		return 0
	}

	if dp[day] >= 0 {
		return dp[day]
	}

	result := math.MaxInt32
	j := day
	for i, duration := range []int{1, 7, 30} {
		for j < len(days) && days[j] < days[day] + duration {
			j++
		}
		result = min(result, helper(j, days, costs, dp) + costs[i])
	}

	dp[day] = result
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 11
	// days := []int{1,4,6,7,8,20}
	// costs := []int{2,7,15}

	// result: 17
	days := []int{1,2,3,4,5,6,7,8,9,10,30,31}
	costs := []int{2,7,15}

	// result: 
	// days := []int{}
	// costs := []int{}

	result := mincostTickets(days, costs)
	fmt.Printf("result = %v\n", result)
}
