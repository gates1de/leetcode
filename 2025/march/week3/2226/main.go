package main
import (
	"fmt"
)

func maximumCandies(candies []int, k int64) int {
	maxCandiesInPile := int(0)
	for _, candy := range candies {
		maxCandiesInPile = max(maxCandiesInPile, candy)
	}

	left := int64(0)
	right := int64(maxCandiesInPile)

	for left < right {
		mid := (left + right + 1) / 2

		if canAllocateCandies(candies, k, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return int(left)
}

func canAllocateCandies(candies []int, k int64, numOfCandies int64) bool {
	maxNumOfChildren := int64(0)

	for _, candy := range candies {
		maxNumOfChildren += int64(candy) / numOfCandies
	}

	return maxNumOfChildren >= k
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// candies := []int{5,8,6}
	// k := int64(3)

	// result: 0
	candies := []int{2,5}
	k := int64(11)

	// result: 
	// candies := []int{}
	// k := int64(0)

	result := maximumCandies(candies, k)
	fmt.Printf("result = %v\n", result)
}

