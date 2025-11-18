package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	i := int(0)
	for i < len(bits) - 1 {
		i += bits[i] + 1
	}

	return i == len(bits) - 1
}

func main() {
	// result: true
	// bits := []int{1,0,0}

	// result: false
	bits := []int{1,1,1,0}

	// result: 
	// bits := []int{}

	result := isOneBitCharacter(bits)
	fmt.Printf("result = %v\n", result)
}

