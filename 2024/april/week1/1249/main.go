package main
import (
	"fmt"
)

func minRemoveToMakeValid(s string) string {
	deletedPos := make([]int, 0)

	for i, char := range s {
		if char == ')' {
			if len(deletedPos) > 0 && s[deletedPos[len(deletedPos) - 1]] == '(' {
				deletedPos = deletedPos[:len(deletedPos) - 1]
			} else {
				deletedPos = append(deletedPos, i)
			}
		} else if char == '(' {
			deletedPos = append(deletedPos, i)
		}
	}

	if len(deletedPos) == 0 {
		return s
	}

	result := make([]rune, 0)
	cursor := 0
	for i, char := range s {
		if cursor < len(deletedPos) && i == deletedPos[cursor] {
			cursor++
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}

func main() {
	// result: "lee(t(c)o)de"
	// s := "lee(t(c)o)de)"

	// result: "ab(c)d"
	// s := "a)b(c)d"

	// result: ""
	s := "))(("

	// result: ""
	// s := ""

	result := minRemoveToMakeValid(s)
	fmt.Printf("result = %v\n", result)
}

