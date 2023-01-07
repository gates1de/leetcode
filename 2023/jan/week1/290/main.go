package main
import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	alphabet := byte('a')
	strs := strings.Split(s, " ")
	m := map[string]byte{}
	convertedS := ""
	for _, str := range strs {
		if m[str] == 0 {
			m[str] = alphabet
			alphabet++
		}
		convertedS += string(m[str])
	}

	alphabet = byte('a')
	patternM := map[rune]byte{}
	convertedPattern := ""
	for _, char := range pattern {
		if patternM[char] == 0 {
			patternM[char] = alphabet
			alphabet++
		}
		convertedPattern += string(patternM[char])
	}

	return convertedPattern == convertedS
}

func main() {
	// result: true
	// pattern := "abba"
	// s := "dog cat cat dog"

	// result: false
	// pattern := "abba"
	// s := "dog cat cat fish"

	// result: false
	// pattern := "aaaa"
	// s := "dog cat cat dog"

	// result: true
	// pattern := "abbca"
	// s := "dog cat cat fish dog"

	// result: true
	pattern := "e"
	s := "eukera"

	// result: 
	// pattern := ""
	// s := ""

	result := wordPattern(pattern, s)
	fmt.Printf("result = %v\n", result)
}

