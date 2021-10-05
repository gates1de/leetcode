package main
import (
	"fmt"
)

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	one, two := 1, 1
	for i := 2; i < n; i++ {
		next := one + two
		one = two
		two = next
	}
	return one + two
}

// Slow & Use more memory
func mySolution(n int) int {
	steps := make([]int, n + 1)
	helper(0, n + 1, steps)
	return steps[n]
}

func helper(i int, n int, steps []int) {
	if i == n {
		return
	}

	if i <= 2 {
		steps[i] = i
	} else {
		steps[i] = steps[i - 2] + steps[i - 1]
	}

	helper(i + 1, n, steps)
}

func main() {
	// result: 2
	// n := int(2)

	// result: 3
	// n := int(3)

	// result: 1
	// n := int(1)

	// result: 89
	// n := int(10)

	// result: 1836311903
	n := int(45)

	// result: 
	// n := int(0)

	result := climbStairs(n)
	fmt.Printf("result = %v\n", result)
}

