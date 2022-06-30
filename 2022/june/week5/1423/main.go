package main
import (
	"fmt"
)

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	sum := int(0)
	for i, point := range cardPoints {
		sum += point
		if i == k - 1 {
			break
		}
	}

	result := sum
	for i := 0; i < k; i++ {
		addIndex := n - i - 1
		removeIndex := addIndex + k
		if removeIndex >= n {
			removeIndex -= n
		}

		sum += cardPoints[addIndex]
		sum -= cardPoints[removeIndex]
		if result < sum {
			result = sum
		}
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

	// result: 9
	cardPoints := []int{1, 2, 3, 4, 5, 6, 2, 1}
	k := int(3)

	// result: 
 	// cardPoints := []int{}
	// k := int(0)

	result := maxScore(cardPoints, k)
	fmt.Printf("result = %v\n", result)
}

