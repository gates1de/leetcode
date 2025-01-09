package main
import (
	"fmt"
)

func prefixCount(words []string, pref string) int {
	result := int(0)
	for _, word := range words {
		if len(pref) > len(word) {
			continue
		}

		if word[:len(pref)] == pref {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// words := []string{"pay","attention","practice","attend"}
	// pref := "at"

	// result: 0
	words := []string{"leetcode","win","loops","success"}
	pref := "code"

	// result: 
	// words := []string{}
	// pref := ""

	result := prefixCount(words, pref)
	fmt.Printf("result = %v\n", result)
}

