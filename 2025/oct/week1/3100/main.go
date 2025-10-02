package main

import (
	"fmt"
	"math"
)

func maxBottlesDrunk(numBottles int, numExchange int) int {
	a := 1
	b := 2 * numExchange - 3
	c := -2 * numBottles
	delta := float64(b * b - 4 * a * c)
	t := int(math.Ceil((float64(-b) + math.Sqrt(delta)) / (2.0 * float64(a))))
	return numBottles + t - 1
}

func main() {
	// result: 15
	// numBottles := int(13)
	// numExchange := int(6)

	// result: 13
	numBottles := int(10)
	numExchange := int(3)

	// result: 
	// numBottles := int(0)
	// numExchange := int(0)

	result := maxBottlesDrunk(numBottles, numExchange)
	fmt.Printf("result = %v\n", result)
}

