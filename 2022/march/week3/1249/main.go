package main
import (
	"fmt"
)

func minRemoveToMakeValid(str string) string {
    s := []byte(str)
    stack := make([]int, 0, len(s))

    for i, v := range s {
        if v == '(' {
            stack = append(stack, i)
        } else if v == ')' {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            } else {
                s[i] = '*'
            }
        }
    }

    for _, i := range stack {
        s[i] = '*'
    }

    res := make([]byte, 0, len(s))
    for i := 0; i < len(s); i++ {
        if s[i] == '*' {
            continue
        }
        res = append(res, s[i])
    }
    return string(res)
}

// Slow & Use more memory
func mySolution(s string) string {
	deleteBraIndices := []int{}
	for i, r := range s {
		n := len(deleteBraIndices)
		if r == rune('(') {
			deleteBraIndices = append(deleteBraIndices, i)
		} else if r == rune(')') {
			if n == 0 {
				s = s[:i] + "*" + s[i + 1:]
			} else {
				deleteBraIndices = deleteBraIndices[:n - 1]
			}
		}
	}

	currentS := s
	for i, index := range deleteBraIndices {
		adjustIndex := index - i
		currentS = currentS[:adjustIndex] + currentS[adjustIndex + 1:]
	}

	result := ""
	for _, r := range currentS {
		if r == rune('*') {
			continue
		}
		result += string(r)
	}

	return result
}

func main() {
	// result: "lee(t(c)o)de"
	// s := "lee(t(c)o)de)"

	// result: "ab(c)d"
	// s := "a)b(c)d"

	// result: ""
	// s := "))(("

	// result: "a(b(c)d)"
	// s := "(a(b(c)d)"

	// result: "()()"
	s := "())()((("

	// result: 
	// s := ""

	result := minRemoveToMakeValid(s)
	fmt.Printf("result = %v\n", result)
}

