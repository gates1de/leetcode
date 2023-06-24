package main
import (
	"fmt"
)

func tallestBillboard(rods []int) int {
	dp := make(map[int]int)
	dp[0] = 0

	for _, rod := range rods {
		newDp := make(map[int]int)
		for key, val := range dp {
			newDp[key] = val
		}

		for diff, taller := range dp {
			shorter := taller - diff

			newTaller := newDp[diff + rod]
			newDp[diff + rod] = max(newTaller, taller + rod)

			newDiff := shorter + rod - taller
			if newDiff < 0 {
				newDiff *= -1
			}
			newTaller2 := max(shorter + rod, taller)
			newDp[newDiff] = max(newTaller2, newDp[newDiff])
		}

		dp = newDp
	}

	return dp[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// rods := []int{1,2,3,6}

	// result: 10
	// rods := []int{1,2,3,4,5,6}

	// result: 0
	rods := []int{1,2}

	// result: 
	// rods := []int{}

	result := tallestBillboard(rods)
	fmt.Printf("result = %v\n", result)
}

