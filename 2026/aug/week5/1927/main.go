package main

import (
	"fmt"
)

func sumGame(num string) bool {
	difference, questionDifference := 0, 0
	half := len(num) / 2
	for i, character := range num {
		sign := int(1)
		if i >= half {
			sign = -1
		}

		if character == '?' {
			questionDifference += sign
		} else {
			difference += sign * int(character - '0')
		}
	}

	if questionDifference % 2 != 0 {
		return true
	}

	return difference != -questionDifference / 2 * 9
}

func main() {
	// result: false
	// num := "5023"

	// result: true
	// num := "25??"

	// result: false
	num := "?3295???"

	// result:
	// num := ""

	result := sumGame(num)
	fmt.Printf("result = %v\n", result)
}
