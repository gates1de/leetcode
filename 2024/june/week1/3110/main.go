package main
import (
	"fmt"
)

func scoreOfString(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := int(0)
	pre := byte(s[0])

	for _, char := range s {
		result += abs(int(pre), int(char))
		pre = byte(char)
	}

	return result
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// result: 13
	// s := "hello"

	// result: 50
	s := "zaz"

	// result: 
	// s := ""

	result := scoreOfString(s)
	fmt.Printf("result = %v\n", result)
}

