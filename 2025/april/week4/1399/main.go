package main
import (
	"fmt"
)

func countLargestGroup(n int) int {
	m := make(map[int]int)
	maxVal := int(0)

	for i := 1; i <= n; i++ {
		key := int(0)
		tmp := i
		for tmp > 0 {
			key += tmp % 10
			tmp /= 10
		}
		m[key]++
		maxVal = max(maxVal, m[key])
	}

	result := int(0)
	for _, val := range m {
		if val == maxVal {
			result++
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
	// result: 4
	// n := int(13)

	// result: 2
	n := int(2)

	// result: 
	// n := int(0)

	result := countLargestGroup(n)
	fmt.Printf("result = %v\n", result)
}

