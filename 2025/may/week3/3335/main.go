package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func lengthAfterTransformations(s string, t int) int {
	m := make([]int, 26)
	for _, char := range s {
		m[char - 'a']++
	}

	for round := 0; round < t; round++ {
		next := make([]int, 26)
		next[0] = m[25]
		next[1] = (m[25] + m[0]) % modulo

		for i := 2; i < 26; i++ {
			next[i] = m[i - 1]
		}

		m = next
	}

	result := int(0)
	for i := 0; i < 26; i++ {
		result = (result + m[i]) % modulo
	}

	return result
}

func main() {
	// result: 7
	// s := "abcyy"
	// t := int(2)

	// result: 5
	s := "azbk"
	t := int(1)

	// result: 
	// s := ""
	// t := int(0)

	result := lengthAfterTransformations(s, t)
	fmt.Printf("result = %v\n", result)
}

