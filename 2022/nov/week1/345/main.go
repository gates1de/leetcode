package main
import (
	"fmt"
)

const LOW_A = byte(97)
const LOW_E = byte(101)
const LOW_I = byte(105)
const LOW_O = byte(111)
const LOW_U = byte(117)

const UP_A = byte(65)
const UP_E = byte(69)
const UP_I = byte(73)
const UP_O = byte(79)
const UP_U = byte(85)

func reverseVowels(s string) string {
	chars := []byte(s)
	i := int(0)
	j := len(chars) - 1
	for i < j {
		left := chars[i]
		right := chars[j]
		isLeftVowel := left == LOW_A || left == LOW_E || left == LOW_I || left == LOW_O || left == LOW_U ||
					  left == UP_A || left == UP_E || left == UP_I || left == UP_O || left == UP_U
		isRightVowel := right == LOW_A || right == LOW_E || right == LOW_I || right == LOW_O || right == LOW_U ||
					  right == UP_A || right == UP_E || right == UP_I || right == UP_O || right == UP_U
		if isLeftVowel && isRightVowel {
			chars[i], chars[j] = chars[j], chars[i]
			i++
			j--
			continue
		}

		if !isLeftVowel {
			i++
		}
		if !isRightVowel {
			j--
		}
	}

	return string(chars)
}

func main() {
	// result: "holle"
	// s := "hello"

	// result: "leotcede"
	// s := "leetcode"

	// result: "Aa"
	s := "aA"

	// result: 
	// s := ""

	result := reverseVowels(s)
	fmt.Printf("result = %v\n", result)
}

