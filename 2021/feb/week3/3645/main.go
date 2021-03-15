package main
import (
	"fmt"
)

// A little faster but use more memory
func minRemoveToMakeValid(s string) string {
	deleteBraIndices := []int{}
	for i, r := range s {
		len := len(deleteBraIndices)
		if r == rune('(') {
			deleteBraIndices = append(deleteBraIndices, i)
		} else if r == rune(')') {
			if len == 0 {
				s = s[:i] + "*" + s[i + 1:]
			} else {
				deleteBraIndices = popInt(deleteBraIndices)
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

// Slow
func myAnswer(s string) string {
	queue := []rune{}
	deleteBraIndices := []int{}
	currentS := ""
	for i, r := range s {
		queueLen := len(queue)
		if r == rune('(') {
			queue = append(queue, r)
			deleteBraIndices = append(deleteBraIndices, i)
		} else if r == rune(')') {
			if queueLen == 0 {
				deleteBraIndices = append(deleteBraIndices, i)
			} else if queue[queueLen - 1] == rune('(') {
				queue = pop(queue)
				deleteBraIndices = popInt(deleteBraIndices)
			}
		}
		currentS += string(r)
	}
	result := s
	for i, index := range deleteBraIndices {
		adjustIndex := index - i
		result = result[:adjustIndex] + result[adjustIndex + 1:]
	}
	return result
}

func pop(queue []rune) []rune {
	return copySlice(queue[:len(queue) - 1])
}

func popInt(queue []int) []int {
	return copyIntSlice(queue[:len(queue) - 1])
}

func copySlice(target []rune) []rune {
    result := make([]rune, len(target))
    copy(result, target)
    return result
}

func copyIntSlice(target []int) []int {
    result := make([]int, len(target))
    copy(result, target)
    return result
}

func main() {
	// s := "lee(t(c)o)de)"
	// s := "a)b(c)d"
	// s := "))(("
	// s := "(a(b(c)d)"
	s := "())()((("

	result := minRemoveToMakeValid(s)
	fmt.Printf("result = %v\n", result)
}

