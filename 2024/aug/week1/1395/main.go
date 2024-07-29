package main
import (
	"fmt"
)

func numTeams(rating []int) int {
	maxRating := int(0)
	for _, rate := range rating {
		maxRating = max(maxRating, rate)
	}

	leftBIT := make([]int, maxRating + 1)
	rightBIT := make([]int, maxRating + 1)

	for _, rate := range rating {
		updateBIT(rightBIT, rate, 1)
	}

	result := int(0)
	for _, rate := range rating {
		updateBIT(rightBIT, rate, -1)

		smallerLeft := getPrefixSum(leftBIT, rate - 1)
		smallerRight := getPrefixSum(rightBIT, rate - 1)

		largerLeft := getPrefixSum(leftBIT, maxRating) - getPrefixSum(leftBIT, rate)
		largerRight := getPrefixSum(rightBIT, maxRating) - getPrefixSum(rightBIT, rate)

		result += smallerLeft * largerRight
		result += smallerRight * largerLeft

		updateBIT(leftBIT, rate, 1)
	}

	return result
}

func updateBIT(bit []int, index int, value int) {
	for index < len(bit) {
		bit[index] += value
		index += index & (-index)
	}
}

func getPrefixSum(bit []int, index int) int {
	sum := int(0)
	for index > 0 {
		sum += bit[index]
		index -= index & (-index)
	}
	return sum
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// rating := []int{2,5,3,4,1}

	// result: 0
	// rating := []int{2,1,3}

	// result: 4
	rating := []int{1,2,3,4}

	// result: 
	// rating := []int{}

	result := numTeams(rating)
	fmt.Printf("result = %v\n", result)
}

