package main
import (
	"fmt"
)

func maxNumberOfBalloons(text string) int {
	m := map[rune]int{}
	m[rune('b')] = 0
	m[rune('a')] = 0
	m[rune('l')] = 0
	m[rune('o')] = 0
	m[rune('n')] = 0
	for _, r := range text {
		if r == rune('b') ||
			r == rune('a') ||
			r == rune('l') ||
			r == rune('o') ||
			r == rune('n') {
			  m[r]++
		}
	}

	result := int(100000)
	for r, count := range m {
		if r == rune('l') || r == rune('o') {
			count /= 2
		}
		if result > count {
			result = count
		}
	}

	return result
}

func main() {
	// result: 1
	// text := "nlaebolko"

	// result: 2
	// text := "loonbalxballpoon"

	// result: 0
	// text := "leetcode"

	// result: 0
	text := "lloo"

	// result: 
	// text := ""

	result := maxNumberOfBalloons(text)
	fmt.Printf("result = %v\n", result)
}

