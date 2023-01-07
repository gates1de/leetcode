package main
import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	totalSurplus := int(0)
	surplus := int(0)
	start := int(0)
	for i := 0; i < n; i++ {
		totalSurplus += gas[i] - cost[i]
		surplus += gas[i] - cost[i]
		if surplus < 0 {
			surplus = 0
			start = i + 1
		}
	}
	if totalSurplus < 0 {
		return -1
	}
	return start
}

func main() {
	// result: 3
	// gas := []int{1,2,3,4,5}
	// cost := []int{3,4,5,1,2}

	// result: -1
	gas := []int{2,3,4}
	cost := []int{3,4,3}

	// result: 
	// gas := []int{}
	// cost := []int{}

	result := canCompleteCircuit(gas, cost)
	fmt.Printf("result = %v\n", result)
}

