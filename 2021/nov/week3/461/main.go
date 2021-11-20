package main
import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	binX := fmt.Sprintf("%032b", x)
	binY := fmt.Sprintf("%032b", y)

	result := int(0)
	for i, d1 := range binX {
		d2 := binY[i]
		if byte(d1) != d2 {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// x := int(1)
	// y := int(4)

	// result: 1
	// x := int(3)
	// y := int(1)

	// result: 18
	x := int(2147483647)
	y := int(1378146394)

	// result: 
	// x := int(0)
	// y := int(0)

	result := hammingDistance(x, y)
	fmt.Printf("result = %v\n", result)
}

