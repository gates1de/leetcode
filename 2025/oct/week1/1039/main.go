package main

import (
	"fmt"
	"math"
)

func minScoreTriangulation(values []int) int {
	m := make(map[int]int)
	return dp(0, len(values) - 1, values, m)
}

func dp(i, j int, values []int, m map[int]int) int {
	n := len(values)

	if i + 2 > j {
		return 0
	}

	if i + 2 > j {
		return values[i] * values[i + 1] * values[j]
	}

	key := i * n + j
	if _, exists := m[key]; !exists {
		minScore := math.MaxInt32
		for k := i + 1; k < j; k++ {
			minScore = min(minScore, values[i] * values[j] * values[k] + dp(i, k, values, m) + dp(k, j, values, m))
		}
		m[key] = minScore
	}

	return m[key]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// values := []int{1,2,3}

	// result: 144
	// values := []int{3,7,4,5}

	// result: 13
	values := []int{1,3,1,4,1,5}

	// result: 
	// values := []int{}

	result := minScoreTriangulation(values)
	fmt.Printf("result = %v\n", result)
}

