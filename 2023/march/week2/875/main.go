package main
import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	right := int(0)
	sumPiles := int(0)
	for _, pile := range piles {
		right = max(right, pile)
		sumPiles += pile
	}

	left := min(1, sumPiles / h)
	for left < right {
		mid := (left + right) / 2
		if speedEnough(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if right == 0 {
		return 1
	}

	return right
}

func speedEnough(piles []int, speed int, hours int) bool {
	actualHours := int(0)
	for _, pile := range piles {
		actualHours += int(math.Ceil(float64(pile) / float64(speed)))
	}
	return actualHours <= hours
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// piles := []int{3,6,7,11}
	// h := int(8)

	// result: 30
	// piles := []int{30,11,23,4,20}
	// h := int(5)

	// result: 23
	piles := []int{30,11,23,4,20}
	h := int(6)

	// result: 
	// piles := []int{}
	// h := int(0)

	result := minEatingSpeed(piles, h)
	fmt.Printf("result = %v\n", result)
}

