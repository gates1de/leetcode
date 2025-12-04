package main

import (
	"fmt"
)

func countCollisions(directions string) int {
	result := int(0)
	count := int(-1)

	for _, dir := range directions {
		if dir == 'R' {
			if count >= 0 {
				count++
			} else {
				count = 1
			}
		} else if dir == 'L' {
			if count >= 0 {
				result += count + 1
				count = 0
			}
		} else if dir == 'S' {
			if count > 0 {
				result += count
			}
			count = 0
		}
	}

	return result
}

func main() {
	// result: 5
	// directions := "RLRSLL"

	// result: 0
	directions := "LLRR"

	// result: 
	// directions := ""

	result := countCollisions(directions)
	fmt.Printf("result = %v\n", result)
}

