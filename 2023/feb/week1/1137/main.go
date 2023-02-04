package main
import (
	"fmt"
)

func tribonacci(n int) int {
    m := map[int]int{0: 0, 1: 1, 2: 1}
    return helper(n, m)
}

func helper(n int, m map[int]int) int {
    if val, ok := m[n]; ok {
        return val
    }
    m[n] = helper(n - 3, m) + helper(n - 2, m) + helper(n - 1, m)
    return m[n]
}

func main() {
	// result: 4
	// n := int(4)

	// result: 1389537
	// n := int(25)

	// result: 2082876103
	n := int(37)

	// result: 
	// n := int(0)

	result := tribonacci(n)
	fmt.Printf("result = %v\n", result)
}

