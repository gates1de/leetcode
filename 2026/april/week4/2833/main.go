package main

import (
	"fmt"
	"math"
)

func furthestDistanceFromOrigin(moves string) int {
	left := int(0)
	right := int(0)
	bar := int(0)
	for _, c := range moves {
		if c == 'L' {
			left++
		} else if c == 'R' {
			right++
		} else {
			bar++
		}
	}

	return int(math.Abs(float64(left - right))) + bar
}

func main() {
	// result: 3
	// moves := "L_RL__R"

	// result: 5
	// moves := "_R__LL_"

	// result: 7
	moves := "_______"

	// result: 
	// moves := ""

	result := furthestDistanceFromOrigin(moves)
	fmt.Printf("result = %v\n", result)
}

