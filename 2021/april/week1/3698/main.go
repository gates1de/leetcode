package main
import (
	"fmt"
)

func minOperations(n int) int {
	if n <= 1 {
		return 0
	}
	if n % 2 == 0 {
		return (n / 2) * (n / 2)
	}
	return (n / 2) * (n / 2 + 1)
}

func mySolution1(n int) int {
	if n <= 1 {
		return 0
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = (2 * i) + 1
	}
	fmt.Printf("nums = %v, n = %v\n", nums, n)
	result := int(0)
	for i := 0; i < n / 2; i++ {
		for nums[i] < n {
			nums[i]++
			result++
		}
	}
	return result
}

func main() {
	// result: 2
	// n := int(3)

	// result: 9
	// n := int(6)

	// result: 16 
	// n := int(8)

	// result: 20
	n := int(9)

	result := minOperations(n)
	fmt.Printf("result = %v\n", result)
}

