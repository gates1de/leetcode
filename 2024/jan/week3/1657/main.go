package main
import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)

	for _, char := range word1 {
		freq1[char - 'a']++
	}
	for _, char := range word2 {
		freq2[char - 'a']++
	}

	for i, _ := range freq1 {
		if (freq1[i] == 0 && freq2[i] != 0) || (freq1[i] != 0 && freq2[i] == 0) {
			return false
		}
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	for i, _ := range freq1 {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// word1 := "abc"
	// word2 := "cba"

	// result: false
	// word1 := "a"
	// word2 := "aa"

	// result: true
	word1 := "cabbba"
	word2 := "abbccc"

	// result: 
	// word1 := ""
	// word2 := ""

	result := closeStrings(word1, word2)
	fmt.Printf("result = %v\n", result)
}

