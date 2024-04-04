package main
import (
	"fmt"
)

func maxDepth(s string) int {
	result := int(0)
	openBrackets := int(0)

	for _, char := range s {
		if char == '(' {
			openBrackets++
		} else if char == ')' {
			openBrackets--
		}
	
		result = max(result, openBrackets)
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// s := "(1+(2*3)+((8)/4))+1"

	// result: 3
	s := "(1)+((2))+(((3)))"

	// result: 
	// s := ""

	result := maxDepth(s)
	fmt.Printf("result = %v\n", result)
}

