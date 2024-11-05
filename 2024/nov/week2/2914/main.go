package main
import (
	"fmt"
)

func minChanges(s string) int {
	result := int(0)

	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i + 1] {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// s := "1001"

	// result: 1
	// s := "10"

	// result: 0
	s := "00"

	// result: 
	// s := ""

	result := minChanges(s)
	fmt.Printf("result = %v\n", result)
}

