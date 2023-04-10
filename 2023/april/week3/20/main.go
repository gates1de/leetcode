package main
import (
	"fmt"
)

func isValid(s string) bool {
	chars := make([]byte, len(s))
	for i, c := range s {
		chars[i] = byte(c)
	}
    charsLen := len(chars)
    if charsLen == 0 || charsLen % 2 == 1 {
        return false
    }

    stack := []byte{}
    for _, c := range chars {
        if c == '(' || c == '{' || c == '[' {
            stack = append(stack, c)
            continue
        }

        if len(stack) == 0 && (c == ')' || c == '}' || c == ']') {
            return false
        }

        target := stack[len(stack) - 1]
        if target == '(' && c == ')' ||
            target == '{' && c == '}' ||
            target == '[' && c == ']' {
            stack = stack[:len(stack) - 1]
        } else {
            return false
        }
    }

    return len(stack) == 0
}

func main() {
	// result: true
	// s := "()"

	// result: true
	// s := "()[]{}"

	// result: false
	s := "(]"

	// result: 
	// s := ""

	result := isValid(s)
	fmt.Printf("result = %v\n", result)
}

