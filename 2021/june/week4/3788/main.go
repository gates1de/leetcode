package main
import (
	"fmt"
)

// REF: https://leetcode.com/problems/number-of-matching-subsequences/discuss/117634/Efficient-and-simple-go-through-words-in-parallel-with-explanation
func numMatchingSubseq(S string, words []string) (num int) {
    waiting := map[rune][]string{' ': words}
    for _, c := range " " + S {
        advance := waiting[c]
        delete(waiting, c)
        for _, word := range advance {
            if len(word) == 0 {
                num++
            } else {
                c := rune(word[0])
                waiting[c] = append(waiting[c], word[1:])
            }
        }
    }
    return
}

// Slow
func mySolution(s string, words []string) int {
	result := int(0)
	for _, word := range words {
		if len(s) < len(word) {
			continue
		}
		if isSubseq(s, word) {
			result++
		}
	}
	return result
}

// w1 = s, w2 = word
func isSubseq(w1 string, w2 string) bool {
	// fmt.Printf("w1 = %v, w2 = %v\n", w1, w2)
	i := int(0)
	j := int(0)
	for i < len(w1) && j < len(w2) {
		c1 := w1[i]
		c2 := w2[j]
		// fmt.Printf("c1 = %v, c2 = %v\n", c1, c2)
		if c1 == c2 {
			i++
			j++
		} else {
			i++
		}
	}

	if j == len(w2) {
		return true
	}
	return false
}

func main() {
	// result: 3
	// s := "abcde"
	// words := []string{"a", "bb", "acd", "ace"}

	// result: 2
	// s := "dsahjpjauf"
	// words := []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}

	// result: 2
	s := "dsahjpjauf"
	words := []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax", "djf"}

	// result: 
	// s := ""
	// words := []string{}

	result := numMatchingSubseq(s, words)
	fmt.Printf("result = %v\n", result)
}

