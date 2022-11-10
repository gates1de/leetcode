package main
import (
	"fmt"
)

func removeDuplicates(s string) string {
	chars := []byte(s)
	stack := make([]byte, 0, len(chars))

	for i := 0; i < len(chars); i++ {
		char := chars[i]
		if len(stack) > 0 {
			last := stack[len(stack) - 1]

			if last == char {
				stack = stack[:len(stack) - 1]
				continue
			}
		}
		stack = append(stack, char)
	}

	return string(stack)
}

func main() {
	// result: "ca"
	// s := "abbaca"

	// result: "ay"
	s := "azxxzy"

	// result: 
	// s := ""

	result := removeDuplicates(s)
	fmt.Printf("result = %v\n", result)
}

