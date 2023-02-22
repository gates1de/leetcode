package main
import (
	"fmt"
)

func shipWithinDays(weights []int, days int) int {
	totalWeight := int(0)
	maxWeight := int(0)
	for _, weight := range weights {
		totalWeight += weight
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	left := maxWeight
	right := totalWeight
	for left < right {
		midWeight := (left + right) / 2
		if feasible(midWeight, days, weights) {
			right = midWeight
		} else {
			left = midWeight + 1
		}
	}
	return left
}

func feasible(midWeight int, days int, weights []int) bool {
	daysNeeded := int(1)
	currentWeight := int(0)
	for _, weight := range weights {
		currentWeight += weight
		if currentWeight > midWeight {
			daysNeeded++
			currentWeight = weight
		}
	}
	return daysNeeded <= days
}

func main() {
	// result: 15
	// weights := []int{1,2,3,4,5,6,7,8,9,10}
	// days := int(5)

	// result: 6
	// weights := []int{3,2,2,4,1,4}
	// days := int(3)

	// result: 3
	weights := []int{1,2,3,1,1}
	days := int(4)

	// result: 
	// weights := []int{}
	// days := int(0)

	result := shipWithinDays(weights, days)
	fmt.Printf("result = %v\n", result)
}

