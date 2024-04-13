package main
import (
	"fmt"
)

func timeRequiredToBuy(tickets []int, k int) int {
	result := int(0)

	for i, ticket := range tickets {
		if i <= k {
			result += min(tickets[k], ticket)
		} else {
			result += min(tickets[k] - 1, ticket)
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// tickets := []int{2,3,2}
	// k := int(2)

	// result: 8
	tickets := []int{5,1,1,1}
	k := int(0)

	// result: 
	// tickets := []int{}
	// k := int(0)

	result := timeRequiredToBuy(tickets, k)
	fmt.Printf("result = %v\n", result)
}

