package main

import (
	"fmt"
	"math/bits"
)

func kthCharacter(k int) byte {
	result := byte(0)

	for k != 1 {
		t := bits.Len(uint(k)) - 1

		if 1<<t == k {
			t--
		}

		k -= 1 << t
		result++
	}

	return byte('a' + result)
}

func main() {
	// result: 'b'
	// k := int(5)

	// result: 'c'
	k := int(10)

	// result: 
	// k := int(0)

	result := kthCharacter(k)
	fmt.Printf("result = %v\n", result)
}

