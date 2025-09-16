package main

import (
	"fmt"
)

func replaceNonCoprimes(nums []int) []int {
	result := make([]int, 0)

	for _, num := range nums {
		result = append(result, num)

		for len(result) > 1 {
			a := result[len(result) - 1]
			b := result[len(result) - 2]
			g := gcd(a, b)

			if g > 1 {
				result = result[:len(result) - 2]
				lcm := a / g * b
				result = append(result, lcm)
			} else {
				break
			}
		}
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func main() {
	// result: [12,7,6]
	// nums := []int{6,4,3,2,7,6,2}

	// result: [2,1,1,3]
	nums := []int{2,2,1,1,3,3,3}

	// result: []
	// nums := []int{}

	result := replaceNonCoprimes(nums)
	fmt.Printf("result = %v\n", result)
}

