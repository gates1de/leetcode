package main

import (
	"fmt"
	"math/bits"
)

func kthCharacter(k int64, operations []int) byte {
	result := byte(0)

	for k != 1 {
		t := bits.Len64(uint64(k)) - 1
		if (1 << t) == k {
			t--
		}

		k -= (1 << t)
		if operations[t] != 0 {
			result++
		}
	}

	return byte('a' + (result % 26))
}

func main() {
	// result: 'a'
	// k := int64(5)
	// operations := []int{0,0,0}

	// result: 'b'
	k := int64(10)
	operations := []int{0,1,0,1}

	// result: ''
	// k := int64(0)
	// operations := []int{}

	result := kthCharacter(k, operations)
	fmt.Printf("result = %v\n", result)
}

