package main
import (
	"fmt"
)

func shortestSuperstring(words []string) string {
	superstrings := []string{}
	for i, word := range words {
		newWords := []string{}
		newWords = append(newWords, words[:i]...)
		newWords = append(newWords, words[i + 1:]...)
		helper(newWords, word, &superstrings)
	}
	// fmt.Printf("superstrings = %v\n", superstrings)

	result := ""
	for _, superstring := range superstrings {
		if len(result) == 0 || len(result) > len(superstring) {
			result = superstring
		}
	}
	return result
}

func helper(words []string, current string, superstrings *[]string) {
	if len(words) == 0 {
		// fmt.Printf("current = %v\n", current)
		*superstrings = append(*superstrings, current)
		return
	}

	for i, word := range words {
		newWords := []string{}
		newWords = append(newWords, words[:i]...)
		newWords = append(newWords, words[i + 1:]...)
		superstring := generateSuperstring(current, word)
		// fmt.Printf("superstring = %v\n", superstring)
		helper(newWords, superstring, superstrings)
	}
}

func generateSuperstring(base string, target string) string {
	result := base + target
	for i, _ := range target {
		if len(base) - 1 < i {
			break
		}
		baseSuffix := base[len(base) - 1 - i:]
		targetPrefix := target[:i + 1]
		// fmt.Printf("baseSuffix = %v, targetPrefix = %v\n", baseSuffix, targetPrefix)
		if baseSuffix == targetPrefix {
			result = base + target[i + 1:]
		}
	}
	return result
}

func main() {
	// result: "alexlovesleetcode"
	// words := []string{"alex", "loves", "leetcode"}

	// result: "gctaagttcatgcatc"
	words := []string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}

	// result: 
	// words := []string{}

	result := shortestSuperstring(words)
	fmt.Printf("result = %v\n", result)
}

