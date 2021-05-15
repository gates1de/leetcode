package main
import (
	"fmt"
)

func maxScore(cardPoints []int, k int) int {
	if k == 0 || len(cardPoints) == 0 {
		return 0
	}
	if len(cardPoints) == 1 {
		return cardPoints[0]
	}

	result := int(0)
	if len(cardPoints) == k {
		for _, num := range cardPoints {
			result += num
		}
		return result
	}

	allPoint := sum(cardPoints)
	i := int(0)
	windowSize := len(cardPoints) - k
	for i <= k {
		sub := cardPoints[i:i + windowSize]
		// fmt.Printf("sub = %v\n", sub)
		point := allPoint - sum(sub)
		if result < point {
			result = point
		}
		i++
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: 12
	// cardPoints := []int{1, 2, 3, 4, 5, 6, 1}
	// k := int(3)

	// result: 4
	// cardPoints := []int{2, 2, 2}
	// k := int(2)

	// result: 55
	// cardPoints := []int{9, 7, 7, 9, 7, 7, 9}
	// k := int(7)

	// result: 1
	// cardPoints := []int{1, 1000, 1}
	// k := int(1)

	// result: 202
	// cardPoints := []int{1, 79, 80, 1, 1, 1, 200, 1}
	// k := int(3)

	// result: 248
	// cardPoints := []int{100, 40, 17, 9, 73, 75}
	// k := int(3)

	// result: 232
	// cardPoints := []int{11, 49, 100, 20, 86, 29, 72}
	// k := int(4)

	// result: 536
	cardPoints := []int{96, 90, 41, 82, 39, 74, 64, 50, 30}
	k := int(8)

	// result: 
	// cardPoints := []int{}
	// k := int(0)

	result := maxScore(cardPoints, k)
	fmt.Printf("result = %v\n", result)
}

