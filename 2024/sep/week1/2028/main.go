package main
import (
	"fmt"
)

func missingRolls(rolls []int, mean int, n int) []int {
	sum := int(0)
	for _, roll := range rolls {
		sum += roll
	}

	missingSum := mean * (n + len(rolls)) - sum
	if missingSum < n || missingSum > 6 * n {
		return []int{}
	}

	result := make([]int, n)
	for i := range result {
		result[i] = missingSum / n
		if i < missingSum % n {
			result[i]++
		}
	}

	return result
}

func main() {
	// result: [6,6]
	// rolls := []int{3, 2, 4, 3}
	// mean := int(4)
	// n := int(2)

	// result: [2,3,2,2]
	// rolls := []int{1,5,6}
	// mean := int(3)
	// n := int(4)

	// result: []
	rolls := []int{1,2,3,4}
	mean := int(6)
	n := int(4)

	// result: []
	// rolls := []int{}
	// mean := int(0)
	// n := int(0)

	result := missingRolls(rolls, mean, n)
	fmt.Printf("result = %v\n", result)
}

