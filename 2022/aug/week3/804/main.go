package main
import (
	"fmt"
)

func uniqueMorseRepresentations(words []string) int {
	letters := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}

	m := map[string]bool{}
	for _, word := range words {
		code := ""
		for _, char := range word {
			code += letters[string(char)]
		}
		m[code] = true
	}
	return len(m)
}

func main() {
	// result: 2
	// words := []string{"gin","zen","gig","msg"}

	// result: 1
	// words := []string{"a"}

	// result: 4
	words := []string{"abcdef","hijklmn","opqrstu","vwxyz","abcdef"}

	// result: 
	// words := []string{}

	result := uniqueMorseRepresentations(words)
	fmt.Printf("result = %v\n", result)
}

