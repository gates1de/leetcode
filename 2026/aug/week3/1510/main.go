package main

import (
	"fmt"
)

func winnerSquareGame(n int) bool {
	if n <= 0 {
		return false
	}

	winning := make([]uint64, n/64+1)

	squares := make([]int, 0, 316)
	for k, square := 1, 1; square <= n; k, square = k+1, square+2*k+1 {
		squares = append(squares, square)
	}

	for stones := 1; stones <= n; stones++ {
		for _, square := range squares {
			if square > stones {
				break
			}

			remaining := stones - square
			if winning[remaining/64]&(uint64(1)<<uint(remaining%64)) == 0 {
				winning[stones/64] |= uint64(1) << uint(stones%64)
				break
			}
		}
	}

	return winning[n/64]&(uint64(1)<<uint(n%64)) != 0
}

func main() {
	// result: true
	// n := int(1)

	// result: false
	// n := int(2)

	// result: true
	// n := int(4)

	// result: true
	// n := int(6)

	// result: true
	// n := int(99856)

	// result: true
	// n := int(100000)

	// result: false
	// n := int(32897)

	// result: true
	// n := int(99999)

	// result: true
	n := int(70)

	// result:
	// n := int(0)

	result := winnerSquareGame(n)
	fmt.Printf("result = %v\n", result)
}
