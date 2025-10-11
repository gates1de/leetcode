package main

import (
	"fmt"
)

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	result := -1 << 31

	for i := n - k; i < n; i++ {
		sum := int(0)
		for j := i; j >= 0; j -= k {
			sum += energy[j]
			result = max(result, sum)
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
	// result: 3
	// energy := []int{5,2,-10,-5,1}
	// k := int(3)

	// result: -1
	energy := []int{-2,-3,-1}
	k := int(2)

	// result: 
	// energy := []int{}
	// k := int(0)

	result := maximumEnergy(energy, k)
	fmt.Printf("result = %v\n", result)
}

