package main

import (
	"fmt"
)

func finalValueAfterOperations(operations []string) int {
	m := map[string]int{"X--": -1, "--X": -1, "X++": 1, "++X": 1}
	result := int(0)
	for _, ope := range operations {
		result += m[ope]
	}
	return result
}

func main() {
	// result: 1
	// operations := []string{"--X","X++","X++"}

	// result: 3
	// operations := []string{"++X","++X","X++"}

	// result: 0
	operations := []string{"X++","++X","--X","X--"}

	// result: 
	// operations := []string{}

	result := finalValueAfterOperations(operations)
	fmt.Printf("result = %v\n", result)
}

