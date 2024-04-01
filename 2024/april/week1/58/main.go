package main
import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	result := int(0)
	count := int(0)

	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		if char == ' ' {
			result = max(result, count)
			if result > 0 {
				break
			}
		} else {
			count++
		}
	}

	result = max(result, count)
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// s := "Hello World"

	// result: 4
	// s := "   fly me   to   the moon  "

	// result: 6
	s := "luffy is still joyboy"

	// result: 
	// s := ""

	result := lengthOfLastWord(s)
	fmt.Printf("result = %v\n", result)
}

