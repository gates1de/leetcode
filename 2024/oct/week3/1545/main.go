package main
import (
	"fmt"
)

func findKthBit(n int, k int) byte {
	positionInSection := k & -k
	isInInvertedPart := (((k / positionInSection) >> 1) & 1) == 1
	originalBitIsOne := (k & 1) == 0

	if isInInvertedPart {
		if originalBitIsOne {
			return '0'
		}
		return '1'
	}

	if originalBitIsOne {
		return '1'
	}
	return '0'
}

func main() {
	// result: "0"
	// n := int(3)
	// k := int(1)

	// result: "1"
	n := int(4)
	k := int(11)

	// result: ""
	// n := int(0)
	// k := int(0)

	result := findKthBit(n, k)
	fmt.Printf("result = %v\n", result)
}

