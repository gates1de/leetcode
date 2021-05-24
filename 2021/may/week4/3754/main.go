package main
import (
	"fmt"
)

func toLowerCase(s string) string {
	result := ""
	for _, r := range s {
		if rune('A') <= r && r <= rune('Z') {
			result += string(r + 32)
		} else {
			result += string(r)
		}
	}
	return result
}

func main() {
	// result: "hello"
	// s := "Hello"

	// result: "here"
	// s := "here"

	// result: "lovely"
	s := "LOVELY"

	// result: 
	// s := ""

	result := toLowerCase(s)
	fmt.Printf("result = %v\n", result)
}

