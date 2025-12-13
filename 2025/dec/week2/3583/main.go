package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func specialTriplets(nums []int) int {
	counter := make(map[int]int)
	partialCounter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	result := int(0)
	for _, num := range nums {
		target := num * 2
		leftCount := partialCounter[target]
		partialCounter[num]++
		rightCount := counter[target] - partialCounter[target]

		result = (result + leftCount * rightCount) % modulo
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{6,3,6}

	// result: 1
	// nums := []int{0,1,0,0}

	// result: 2
	nums := []int{8,4,2,8,4}

	// result: 
	// nums := []int{}

	result := specialTriplets(nums)
	fmt.Printf("result = %v\n", result)
}

