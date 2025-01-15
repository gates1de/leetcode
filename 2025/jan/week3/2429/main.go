package main
import (
	"fmt"
	"math/bits"
)

func minimizeXor(num1 int, num2 int) int {
	result := int(0)
	targetSetBitsCount := bits.OnesCount(uint(num2))
	setBitsCount := int(0)
	currentBit := int(31)

	for setBitsCount < targetSetBitsCount {
		if isSet(num1, currentBit) || (targetSetBitsCount - setBitsCount > currentBit) {
			result = setBit(result, currentBit)
			setBitsCount++
		}

		currentBit--
	}

	return result
}

func isSet(x int, bit int) bool {
	return (x & (1 << bit)) != 0
}

func setBit(x int, bit int) int {
	return x | (1 << bit)
}

func main() {
	// result: 3
	// num1 := int(3)
	// num2 := int(5)

	// result: 3
	num1 := int(1)
	num2 := int(12)

	// result: 
	// num1 := int(0)
	// num2 := int(0)

	result := minimizeXor(num1, num2)
	fmt.Printf("result = %v\n", result)
}

