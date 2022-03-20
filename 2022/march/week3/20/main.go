package main
import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	chars := strings.Split(s, "")
	charsLen := len(chars)
	if charsLen == 0 || charsLen % 2 == 1 {
		return false
	}

	stack := []string{}
	for _, c := range chars {
		if strings.Contains("({[", c) {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 && strings.Contains(")}]", c) {
			return false
		}

		target := stack[len(stack) - 1]
		if target == "(" && c == ")" ||
			target == "{" && c == "}" ||
			target == "[" && c == "]" {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}

	return strings.Join(stack, "") == ""
}

func main() {
	// result: true
	// s := "()"

	// result: true
	// s := "()[]{}"

	// result: false
	// s := "(]"

	// result: false
	// s := "([)]"

	// result: true
	// s := "{[]}"

	// result: false
	// s := "(("

	// result: false
	// s := "(]"

	// result: false
	// s := "([}}])"

	// result: false
	s := "){"

	// result: 
	// s := ""

	result := isValid(s)
	fmt.Printf("result = %v\n", result)
}

