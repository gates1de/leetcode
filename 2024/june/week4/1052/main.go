package main
import (
	"fmt"
)

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	unrealized := int(0)
	for i := 0; i < minutes; i++ {
		unrealized += customers[i] * grumpy[i]
	}

	maxUnrealized := unrealized

	for i := minutes; i < n; i++ {
		unrealized += customers[i] * grumpy[i]
		unrealized -= customers[i - minutes] * grumpy[i - minutes]

		maxUnrealized = max(maxUnrealized, unrealized)
	}

	result := maxUnrealized
	for i, customer := range customers {
		result += customer * (1 - grumpy[i])
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 16
	// customers := []int{1,0,1,2,1,1,7,5}
	// grumpy := []int{0,1,0,1,0,1,0,1}
	// minutes := int(3)

	// result: 1
	customers := []int{1}
	grumpy := []int{0}
	minutes := int(1)

	// result: 
	// customers := []int{}
	// grumpy := []int{}
	// minutes := int(0)

	result := maxSatisfied(customers, grumpy, minutes)
	fmt.Printf("result = %v\n", result)
}

