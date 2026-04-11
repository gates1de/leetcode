package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	x := int(0)
	y := int(0)
	for _, char := range moves {
		switch char {
		case 'U':
			y--
		case 'D':
			y++
		case 'L':
			x--
		case 'R':
			x++
		}
	}

	return x == 0 && y == 0
}

func main() {
	// result: true
	// moves := "UD"

	// result: false
	moves := "LL"

	// result: false
	// moves := ""

	result := judgeCircle(moves)
	fmt.Printf("result = %v\n", result)
}

