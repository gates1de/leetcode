package main

import (
	"fmt"
)

func flowerGame(n int, m int) int64 {
	return int64(m) * int64(n) / 2
}

func main() {
	// result: 3
	// m := int(3)
	// n := int(2)

	// result: 0
	m := int(1)
	n := int(1)

	// result: 
	// m := int(0)
	// n := int(0)

	result := flowerGame(n, m)
	fmt.Printf("result = %v\n", result)
}

