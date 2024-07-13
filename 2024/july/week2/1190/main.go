package main
import (
	"fmt"
)

func reverseParentheses(s string) string {
	stack := make([]int, 0, len(s))
	result := make([]byte, 0, len(s))

	for _, char := range s {
		if char == '(' {
			stack = append(stack, len(result))
		} else if char == ')' {
			start := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			reverse(result, start, len(result) - 1)
		} else {
			result = append(result, byte(char))
		}
	}

	return string(result)
}

func reverse(chars []byte, start int, end int) {
	for start < end {
		temp := chars[start]
		chars[start] = chars[end]
		chars[end] = temp
		start++
		end--
	}
}

func main() {
	// result: "dcba"
	// s := "(abcd)"

	// result: "iloveu"
	// s := "(u(love)i)"

	// result: "leetcode"
	s := "(ed(et(oc))el)"

	// result: ""
	// s := ""

	result := reverseParentheses(s)
	fmt.Printf("result = %v\n", result)
}

