package main
import (
	"fmt"
)

func partitionString(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := make(map[byte]bool, 26)
	result := int(1)

	for _, char := range s {
		b := byte(char)
		if m[b] {
			m = make(map[byte]bool, 26)
			m[b] = true
			result++
		} else {
			m[b] = true
		}
	}

	return result
}

func main() {
	// result: 4
	// s := "abacaba"

	// result: 6
	s := "ssssss"

	// result: 
	// s := ""

	result := partitionString(s)
	fmt.Printf("result = %v\n", result)
}

