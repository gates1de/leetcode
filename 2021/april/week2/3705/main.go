package main
import (
	"fmt"
)

func constructArray(n int, k int) []int {
	result := make([]int, n)
	result[0] = 1

	targetDiff := k
	i := int(1)
	if k == 1 {
		for i < n {
			result[i] = i + 1
			i++
		}
		return result
	}

	currentMax := result[0]
	for i < n {
		if i == 1 {
			result[i] = result[i - 1] + k
			targetDiff--
			if currentMax < result[i] {
				currentMax = result[i]
			}
			i++
			continue
		}

		if targetDiff == 0 {
			result[i] = currentMax + 1
			i++
			currentMax++
			continue
		}

		if result[i - 1] - result[i - 2] < 0 {
			result[i] = result[i - 2] - 1
		} else {
			result[i] = result[i - 2] + 1
		}

		if currentMax < result[i] {
			currentMax = result[i]
		}
		i++
		targetDiff--
	}
	return result
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: [1, 2, 3]
	// n := int(3)
	// k := int(1)

	// result: [1, 3, 2]
	// n := int(3)
	// k := int(2)

	// result: [1, 4, 2, 3]
	// n := int(4)
	// k := int(3)

	// result: [1, 5, 2, 4, 3]
	// n := int(5)
	// k := int(4)

	// result: [1, 5, 2, 4, 3, 6]
	n := int(7)
	k := int(4)

	// result: 
	// n := int()
	// k := int()

	result := constructArray(n, k)
	fmt.Printf("result = %v\n", result)
}

