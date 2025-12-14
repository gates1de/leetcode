package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numberOfWays(corridor string) int {
	zero := int(0)
	one := int(0)
	two := int(1)

	for _, char := range corridor {
		if char == 'S' {
			zero = one
			one, two = two, one
		} else {
			two = (two + zero) % modulo
		}
	}

	return zero
}

func main() {
	// result: 3
	// corridor := "SSPPSPS"

	// result: 1
	// corridor := "PPSPSP"

	// result: 0
	corridor := "S"

	// result: 
	// corridor := ""

	result := numberOfWays(corridor)
	fmt.Printf("result = %v\n", result)
}

