package main
import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	m := map[rune]int{}
	for _, char1 := range s1 {
		m[char1]++
	}

	for i, char2 := range s2 {
		if i >= len(s1) {
			removeChar := rune(s2[i - len(s1)])
			m[removeChar]++
			if m[removeChar] == 0 {
				delete(m, removeChar)
			}
		}

		m[char2]--
		if m[char2] == 0 {
			delete(m, char2)
		}

		if len(m) == 0 {
			return true
		}
	}

	return false
}

func main() {
	// result: true
	// s1 := "ab"
	// s2 := "eidbaooo"

	// result: false
	// s1 := "ab"
	// s2 := "eidboaoo"

	// result: false
	// s1 := "a"
	// s2 := "b"

	// result: true
	// s1 := "a"
	// s2 := "a"

	// result: false
	// s1 := "ac"
	// s2 := "abcdefg"

	// result: true
	// s1 := "xyz"
	// s2 := "hijklmnzyxwvopqrstuabcdefg"

	// result: true
	// s1 := "adc"
	// s2 := "dcda"

	// result: true
	// s1 := "adc"
	// s2 := "adbdcda"

	// result: false
	s1 := "adcc"
	s2 := "abcdcda"

	// result: 
	// s1 := ""
	// s2 := ""

	result := checkInclusion(s1, s2)
	fmt.Printf("result = %v\n", result)
}

