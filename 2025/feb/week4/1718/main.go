package main
import (
	"fmt"
)

func constructDistancedSequence(n int) []int {
	result := make([]int, n * 2 - 1)
	isUsed := make([]bool, n + 1)
	findLexicographicallyLargestSequence(0, result, isUsed, n)
	return result
}

func findLexicographicallyLargestSequence(index int, result []int, isUsed []bool, target int) bool {
	if index == len(result) {
		return true
	}

	if result[index] != 0 {
		return findLexicographicallyLargestSequence(index + 1, result, isUsed, target)
	}

	for i := target; i >= 1; i-- {
		if isUsed[i] {
			continue
		}

		isUsed[i] = true
		result[index] = i

		if i == 1 {
			if findLexicographicallyLargestSequence(index + 1, result, isUsed, target) {
				return true
			}
		} else if index + i < len(result) && result[index + i] == 0 {
			result[index + i] = i

			if findLexicographicallyLargestSequence(index + 1, result, isUsed, target) {
				return true
			}

			result[index + i] = 0
		}

		result[index] = 0
		isUsed[i] = false
	}

	return false
}

func main() {
	// result: [3,1,2,3,2]
	// n := int(3)

	// result: [5,3,1,4,3,5,2,4,2]
	n := int(5)

	// result: []
	// n := int(0)

	result := constructDistancedSequence(n)
	fmt.Printf("result = %v\n", result)
}

