package main

import (
	"fmt"
	"math"
)

func findClosest(x int, y int, z int) int {
	dxz := int(math.Abs(float64(x - z)))
	dyz := int(math.Abs(float64(y - z)))

	if dxz < dyz {
		return 1
	} else if dxz > dyz {
		return 2
	}

	return 0
}

func main() {
	// result: 1
	// x := int(2)
	// y := int(7)
	// z := int(4)

	// result: 2
	// x := int(2)
	// y := int(5)
	// z := int(6)

	// result: 0
	x := int(1)
	y := int(5)
	z := int(3)

	// result: 
	// x := int(0)
	// y := int(0)
	// z := int(0)

	result := findClosest(x, y, z)
	fmt.Printf("result = %v\n", result)
}

