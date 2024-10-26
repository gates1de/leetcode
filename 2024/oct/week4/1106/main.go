package main
import (
	"fmt"
)

func parseBoolExpr(expression string) bool {
	stack := make([]rune, 0)

	for _, char := range expression {
		if char == ',' || char == '(' {
			continue
		}

		if char == 't' || char == 'f' || char == '!' || char == '&' || char == '|' {
			stack = append(stack, char)
		} else if char == ')' {
			hasTrue := false
			hasFalse := false

			for len(stack) > 0 && peek(stack) != '!' && peek(stack) != '&' && peek(stack) != '|' {
				top := peek(stack)
				stack = stack[:len(stack) - 1]

				if top == 't' {
					hasTrue = true
				}
				if top == 'f' {
					hasFalse = true
				}
			}

			if len(stack) > 0 {
				operator := peek(stack)
				stack = stack[:len(stack) - 1]

				if operator == '!' {
					if hasTrue {
						stack = append(stack, 'f')
					} else {
						stack = append(stack, 't')
					}
				} else if operator == '&' {
					if hasFalse {
						stack = append(stack, 'f')
					} else {
						stack = append(stack, 't')
					}
				} else {
					if hasTrue {
						stack = append(stack, 't')
					} else {
						stack = append(stack, 'f')
					}
				}
			}
		}
	}

	return peek(stack) == 't'
}

func peek(stack []rune) rune {
	if len(stack) == 0 {
		return 0
	}
	return stack[len(stack) - 1]
}

func main() {
	// result: false
	// expression := "&(|(f))"

	// result: true
	// expression := "|(f,f,f,t)"

	// result: true
	expression := "!(&(f,t))"

	// result: 
	// expression := ""

	result := parseBoolExpr(expression)
	fmt.Printf("result = %v\n", result)
}

