package main
import (
	"fmt"
	"regexp"
)

func numMatchingSubseq(s string, words []string) int {
    waiting := map[rune][]string{' ': words}
	result := int(0)
    for _, c := range " " + s {
        advance := waiting[c]
        delete(waiting, c)
        for _, word := range advance {
            if len(word) == 0 {
                result++
            } else {
                c := rune(word[0])
                waiting[c] = append(waiting[c], word[1:])
            }
        }
    }

    return result
}

// Time Limit Exceeded
func ngSolution(s string, words []string) int {
	result := int(0)
	for _, word := range words {
		if len(s) < len(word) {
			continue
		}

		regexWord := ""
		symbol := ""
		for _, c := range word {
			regexWord = fmt.Sprintf("%v%v%v", regexWord, symbol, string(c))
			symbol = "(.*)"
		}

		regex := regexp.MustCompile(regexWord)
		if regex.MatchString(s) {
			result++
		}
	}

	return result
}

func main() {
	// result: 3
	// s := "abcde"
	// words := []string{"a","bb","acd","ace"}

	// result: 2
	s := "dsahjpjauf"
	words := []string{"ahjpjau","ja","ahbwzgqnuk","tnmlanowax"}

	// result: 
	// s := ""
	// words := []string{}

	result := numMatchingSubseq(s, words)
	fmt.Printf("result = %v\n", result)
}

