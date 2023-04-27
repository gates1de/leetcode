package main
import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func main() {
	// result: 1
	// n := int(3)

	// result: 0
	// n := int(0)

	// result: 1
	n := int(1)

	// result: 
	// n := int(0)

	result := bulbSwitch(n)
	fmt.Printf("result = %v\n", result)
}

