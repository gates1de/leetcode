package main

import (
	"fmt"
)

func sumZero(n int) []int {
	result := make([]int, 0, n)

	for i := 1; i <= n / 2; i++ {
		result = append(result, i, -i)
	}
	if n % 2 == 1 {
		result = append(result, 0)
	}

	return result
}

func main() {
	// result: [-7,-1,1,3,4]
	// n := int(5)

	// result: [-1,0,1]
	// n := int(3)

	// result: [0]
	n := int(1)

	// result: []
	// n := int(0)

	result := sumZero(n)
	fmt.Printf("result = %v\n", result)
}

