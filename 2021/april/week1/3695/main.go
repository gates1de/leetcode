package main
import (
	"fmt"
)

func longestValidParentheses(s string) int {
    max := 0
    l := 0
    r := 0

	// traversing the string from left towards right
    for i := 0; i < len(s); i++ {
        if s[i] == '(' { l++ }
        if s[i] == ')' { r++ }
        if r > l {
            r = 0
            l = 0
        } else if r == l {
            if r + l > max {
                max = r + l
            }
        }
    }

    l = 0
    r = 0
	// traversing the string from right towards left
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ')' { l++ }
        if s[i] == '(' { r++ }
        if r > l {
            r = 0
            l = 0
        } else if r == l {
            if r + l > max {
                max = r + l
            }
        }
    }

    return max
}


// Slow & Use more memory
func mySolution(s string) int {
	deleteBraIndices := []int{}
	for i, r := range s {
		len := len(deleteBraIndices)
		if r == rune('(') {
			deleteBraIndices = append(deleteBraIndices, i)
		} else if r == rune(')') {
			if len == 0 {
				s = s[:i] + " " + s[i + 1:]
			} else {
				deleteBraIndices = pop(deleteBraIndices)
			}
		}
	}

	currentS := s
	for _, index := range deleteBraIndices {
		currentS = currentS[:index] + " " + currentS[index + 1:]
	}
	// fmt.Printf("currentS = %v\n", currentS)
	result := int(0)
	current := int(0)
	for _, r := range currentS {
		if r == rune(' ') {
			current = 0
			continue
		}
		if r == rune('(') {
			current += 2
		}
		if current > result {
			result = current
		}
	}
	return result
}

func copySlice(target []int) []int {
    result := make([]int, len(target))
    copy(result, target)
    return result
}

func pop(queue []int) []int {
	return copySlice(queue[:len(queue) - 1])
}

func main() {
	// result: 2
	// s := "(()"

	// result: 4
	// s := ")()())"

	// result: 2
	// s := "()(()"

	// result: 18
	// s := "()(()())((()()())))"

	// result: 6
	// s := "()(())"

	// result: 6
	// s := ")(((((()())()()))()(()))("

	// result: 2
	s := "))))((()(("

	result := longestValidParentheses(s)
	fmt.Printf("result = %v\n", result)
}

