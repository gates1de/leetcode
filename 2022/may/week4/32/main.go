package main
import (
	"fmt"
)

const LEFT = rune('(')
const RIGHT = rune(')')
const SPACE = rune(' ')

func longestValidParentheses(s string) int {
    deleteBraIndices := []int{}
    for i, r := range s {
        length := len(deleteBraIndices)

        if r == LEFT {
            deleteBraIndices = append(deleteBraIndices, i)
        } else if r == RIGHT {
            if length == 0 {
                s = s[:i] + " " + s[i + 1:]
            } else {
                deleteBraIndices = pop(deleteBraIndices)
            }
        }
    }

    currentS := s
    for _, i := range deleteBraIndices {
        currentS = currentS[:i] + " " + currentS[i + 1:]
    }

    result := int(0)
    current := int(0)
    for _, r := range currentS {
        if r == SPACE {
            current = 0
            continue
        }
        if r == LEFT {
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

	// result: 0
	// s := ""

	// result: 2
	// s := "()(()"

	// result: 6
	// s := "()(())"

	// result: 22
	// s := ")(((((()())()()))()(()))("

	// result: 2
	// s := "))))((()(("

	// result: 4
	// s := ")()())()()("

	// result: 2
	// s := "(()(((()"

	// result: 4
	s := "(()()"

	// result: 
	// s := ""

	result := longestValidParentheses(s)
	fmt.Printf("result = %v\n", result)
}

