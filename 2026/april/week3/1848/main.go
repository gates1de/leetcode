package main

import (
	"fmt"
	"math"
)

func getMinDistance(nums []int, target int, start int) int {
	result := len(nums)
	for i, num := range nums {
		if num == target {
			distance := int(math.Abs(float64(i - start)))
			result = min(result, distance)
		}
	}

	return result
}

func main() {
	// result: 1
	// nums := []int{1,2,3,4,5}
	// target := int(5)
	// start := int(3)

	// result: 0
	// nums := []int{1}
	// target := int(1)
	// start := int(0)

	// result: 
	nums := []int{1,1,1,1,1,1,1,1,1,1}
	target := int(1)
	start := int(0)

	// result: 
	// nums := []int{}
	// target := int(0)
	// start := int(0)

	result := getMinDistance(nums, target, start)
	fmt.Printf("result = %v\n", result)
}

