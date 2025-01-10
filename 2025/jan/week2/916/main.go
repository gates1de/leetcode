package main
import (
	"fmt"
	"strings"
)

func wordSubsets(A []string, B []string) []string {
	BFreq := map[string]int{}
	result := []string{}
	countMax := int(0)

	for _, word := range B {
			for _, r := range word {
					char := string(r)
					count := strings.Count(word, char)

					if BFreq[char] != 0 {
							diff := count - BFreq[char]

							if diff > 0 {
									BFreq[char] = count
									countMax += diff
							}
					} else {
							BFreq[char] = count
							countMax += count
					}
			}

			if countMax > 10 {
					return result
			}
	}

	TOP:
	for _, word := range A {
			if len(word) < countMax {
					continue
			}

			for char, _ := range BFreq {
					if strings.Count(word, char) < BFreq[char] {
							continue TOP
					}
			}
			result = append(result, word)
	}

	return result
}

func main() {
		// result: ["facebook","google","leetcode"]
	// words1 := []string{"amazon","apple","facebook","google","leetcode"}
	// words2 := []string{"e","o"}

	// result: ["apple","google","leetcode"]
	// words1 := []string{"amazon","apple","facebook","google","leetcode"}
	// words2 := []string{"l","e"}

	// result: ["facebook","leetcode"]
	// words1 := []string{"amazon","apple","facebook","google","leetcode"}
	// words2 := []string{"ec","o"}

	// result: ["leetcode"]
	words1 := []string{"amazon","apple","facebook","google","leetcode"}
	words2 := []string{"ece","o"}

	// result: []
	// words1 := []string{}
	// words2 := []string{}

	result := wordSubsets(words1, words2)
	fmt.Printf("result = %v\n", result)
}

