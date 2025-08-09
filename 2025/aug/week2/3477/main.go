package main

import (
	"fmt"
)

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(baskets)
	result := int(0)

	for _, fruit := range fruits {
		unset := int(1)

		for i := 0; i < n; i++ {
			if fruit <= baskets[i] {
				baskets[i] = 0
				unset = 0
				break
			}
		}

		result += unset
	}

	return result
}

func main() {
	// result: 1
	// fruits := []int{4,2,5}
	// baskets := []int{3,5,4}

	// result: 0
	fruits := []int{3,6,1}
	baskets := []int{6,4,7}

	// result: 
	// fruits := []int{}
	// baskets := []int{}

	result := numOfUnplacedFruits(fruits, baskets)
	fmt.Printf("result = %v\n", result)
}

