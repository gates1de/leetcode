package main
import (
	"fmt"
)

func scoreOfParentheses(S string) int {
	result, _ := counter(S, 0)
	return result
}

func counter(S string, currentIndex int) (int, int) {
	// fmt.Printf("S = %v\n", S)
	queue := []rune{}
	count := int(0)
	for i := currentIndex; i < len(S); i++ {
		r := rune(S[i])
		// fmt.Printf("queue = %v, S = %v\n", queue, S[i:])
		if r == rune('(') {
			if len(queue) == 0 {
				queue = append(queue, r)
				continue
			}
			innerCount, newIndex := counter(S, i)
			count += 2 * innerCount
			// fmt.Printf("count = %v, index = %v\n", count, newIndex)
			if newIndex == len(S) - 1 {
				return count, 0
			}
			i = newIndex
		} else if r == rune(')') {
			if len(queue) == 0 {
				return count, i
			}
			if queue[len(queue) - 1] == '(' {
				count++
			}
		}
		queue = []rune{}
	}
	return count, len(S)
}

func main() {
	// S := "()"
	// S := "(())"
	// S := "()()"
	// S := "(()(()))"
	S := "(())()"

	result := scoreOfParentheses(S)
	fmt.Printf("result = %v\n", result)
}

