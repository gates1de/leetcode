package main
import (
	"fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxCandy := int(-1)
	for _, candy := range candies {
		maxCandy = max(maxCandy, candy)
	}

	result := make([]bool, n)
	for i, candy := range candies {
		if maxCandy <= candy + extraCandies {
			result[i] = true
		}
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
	// result: [true,true,true,false,true]
	// candies := []int{2,3,5,1,3}
	// extraCandies := int(3)

	// result: [true,false,false,false,false]
	// candies := []int{4,2,1,1,2}
	// extraCandies := int(1)

	// result: [true,false,true]
	candies := []int{12,1,12}
	extraCandies := int(10)

	// result: 
	// candies := []int{}
	// extraCandies := int(0)

	result := kidsWithCandies(candies, extraCandies)
	fmt.Printf("result = %v\n", result)
}

