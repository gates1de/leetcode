package main

import (
	"fmt"
)

func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	result := int64(need1)*int64(cost1) + int64(need2)*int64(cost2)
	common := min(need1, need2)
	candidate := int64(common)*int64(costBoth) + int64(need1-common)*int64(cost1) + int64(need2-common)*int64(cost2)
	result = min(result, candidate)

	maximum := max(need1, need2)
	candidate = int64(maximum) * int64(costBoth)
	result = min(result, candidate)

	return result
}

func main() {
	// result: 3
	// cost1 := 3
	// cost2 := 2
	// costBoth := 1
	// need1 := 3
	// need2 := 2

	// result: 22
	// cost1 := 5
	// cost2 := 4
	// costBoth := 15
	// need1 := 2
	// need2 := 3

	// result: 0
	cost1 := 5
	cost2 := 4
	costBoth := 15
	need1 := 0
	need2 := 0

	// result: 0
	// cost1 := 0
	// cost2 := 0
	// costBoth := 0
	// need1 := 0
	// need2 := 0

	result := minimumCost(cost1, cost2, costBoth, need1, need2)
	fmt.Printf("result = %v\n", result)
}
