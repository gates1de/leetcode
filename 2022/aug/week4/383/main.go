package main
import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	m := map[rune]int{}
	for _, char := range magazine {
		m[char]++
	}
	for _, char := range ransomNote {
		m[char]--
		if m[char] < 0 {
			return false
		}
	}

	return true
}

func main() {
	// result: false
	// ransomNote := "a"
	// magazine := "b"

	// result: false
	// ransomNote := "aa"
	// magazine := "ab"

	// result: true
	// ransomNote := "aa"
	// magazine := "aab"

	// result: true
	ransomNote := "iewa"
	magazine := "mfoenwalknfewkafewaia"

	// result: 
	// ransomNote := ""
	// magazine := ""

	result := canConstruct(ransomNote, magazine)
	fmt.Printf("result = %v\n", result)
}

