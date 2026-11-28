package main

import (
	"fmt"
)

func canAliceWin(n int) bool {
	aliceTurn := true
	for remove := 10; n >= remove; remove-- {
		n -= remove
		aliceTurn = !aliceTurn
	}

	return !aliceTurn
}

func main() {
	// result: true
	// n := 12

	// result: false
	n := 1

	// result: false
	// n := 0

	result := canAliceWin(n)
	fmt.Printf("result = %v\n", result)
}
