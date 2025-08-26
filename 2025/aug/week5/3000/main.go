package main

import (
	"fmt"
)

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxSqrt := int(0)
	result := int(0)

	for _, dim := range dimensions {
		l := dim[0]
		w := dim[1]
		diaSqrt := l * l + w * w
		area := l * w

		if diaSqrt > maxSqrt {
			maxSqrt = diaSqrt
			result = area
		} else if diaSqrt == maxSqrt {
			if area > result {
				result = area
			}
		}
	}

	return result
}

func main() {
	// result: 48
	// dimensions := [][]int{{9,3},{8,6}}

	// result: 12
	dimensions := [][]int{{3,4},{4,3}}

	// result: 
	// dimensions := [][]int{}

	result := areaOfMaxDiagonal(dimensions)
	fmt.Printf("result = %v\n", result)
}

