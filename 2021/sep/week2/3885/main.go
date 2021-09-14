package main
import (
	"fmt"
)

func reverseOnlyLetters(s string) string {
	i := int(0)
	j := len(s) - 1
	result := make([]rune, len(s))
	for i <= j {
		head := s[i]
		tail := s[j]
		// fmt.Printf("head = %v, tail = %v\n", head, tail)
		if !isLetter(head) {
			result[i] = rune(head)
			i++
			continue
		}
		if !isLetter(tail) {
			result[j] = rune(tail)
			j--
			continue
		}
		result[i] = rune(tail)
		result[j] = rune(head)
		i++
		j--
	}
	return string(result)
}

func isLetter(b byte) bool {
	return (65 <= b && b <= 90) || (97 <= b && b <= 122)
}

func main() {
	// result: "dc-ba"
	// s := "ab-cd"

	// result: "j-Ih-gfE-dCba"
	// s := "a-bC-dEf-ghIj"

	// result: "Qedo1ct-eeLg=ntse-T!"
	// s := "Test1ng-Leet=code-Q!"

	// result: "I"
	s := "I"

	// result: 
	// s := ""

	result := reverseOnlyLetters(s)
	fmt.Printf("result = %v\n", result)
}

