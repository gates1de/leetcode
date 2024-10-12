package main
import (
	"fmt"
)

func minAddToMakeValid(s string) int {
	openBrackets := int(0)
	minAddsRequired := int(0)

	for _, char := range s {
		if char == '(' {
			openBrackets++
		} else {
			if openBrackets > 0 {
				openBrackets--
			} else {
				minAddsRequired++
			}
		}
	}

	return minAddsRequired + openBrackets
}

func main() {
	// result: 1
	// s := "())"

	// result: 3
	s := "((("

	// result: 
	// s := ""

	result := minAddToMakeValid(s)
	fmt.Printf("result = %v\n", result)
}

