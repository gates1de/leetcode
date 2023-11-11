package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countHomogenous(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	result := int(0)
	pre := s[0]
	streak := int(1)

	for _, char := range s[1:] {
		b := byte(char)
		result = (result + streak) % modulo

		if pre != b {
			streak = 1
			pre = b
			continue
		}

		streak++
	}

	result = (result + streak) % modulo
	return result
}

func main() {
	// result: 13
	// s := "abbcccaa"

	// result: 2
	// s := "xy"

	// result: 15
	s := "zzzzz"

	// result: 
	// s := ""

	result := countHomogenous(s)
	fmt.Printf("result = %v\n", result)
}

