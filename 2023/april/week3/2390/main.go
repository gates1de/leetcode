package main
import (
	"fmt"
)

func removeStars(s string) string {
	stack := make([]byte, 0, len(s))
	for _, char := range s {
		b := byte(char)

		if len(stack) == 0 {
			stack = append(stack, b)
			continue
		}

		if b == '*' {
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, b)
		}
	}

	return string(stack)
}

func main() {
	// result: "lecoe"
	// s := "leet**cod*e"

	// result: ""
	s := "erase*****"

	// result: 
	// s := ""

	result := removeStars(s)
	fmt.Printf("result = %v\n", result)
}

