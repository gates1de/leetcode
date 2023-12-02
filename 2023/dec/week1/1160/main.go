package main
import (
	"fmt"
)

func countCharacters(words []string, chars string) int {
	counts := make([]int, 26)
	for _, char := range chars {
		counts[char - 'a']++
	}

	result := int(0)
	for _, word := range words {
		wordCount := make([]int, 26)
		for _, char := range word {
			wordCount[char - 'a']++
		}

		isOk := true
		for i := 0; i < 26; i++ {
			if counts[i] < wordCount[i] {
				isOk = false
				break
			}
		}

		if isOk {
			result += len(word)
		}
	}

	return result
}

func main() {
	// result: 6
	// words := []string{"cat","bt","hat","tree"}
	// chars := "atach"

	// result: 10
	words := []string{"hello","world","leetcode"}
	chars := "welldonehoneyr"

	// result: 
	// words := []string{}
	// chars := ""

	result := countCharacters(words, chars)
	fmt.Printf("result = %v\n", result)
}

