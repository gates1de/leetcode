package main
import (
	"fmt"
)

func maxScoreWords(words []string, letters []byte, score []int) int {
	n := len(words)
	freq := [26]int{}
	for _, c := range letters {
		freq[c - 'a']++
	}

	result := int(0)
	check(n - 1, words, score, freq, [26]int{}, 0, &result)
	return result
}

func check(n int, words []string, score []int, freq [26]int, subsetLetters [26]int, totalScore int, result *int) {
	if n == -1 {
		*result = max(*result, totalScore)
		return
	}

	check(n - 1, words, score, freq, subsetLetters, totalScore, result)

	l := len(words[n])
	for i := 0; i < l; i++ {
		char := words[n][i] - 'a'
		subsetLetters[char]++
		totalScore += score[char]
	}

	if isValidWord(freq, subsetLetters) {
		check(n - 1, words, score, freq, subsetLetters, totalScore, result)
	}

	for i := 0; i < l; i++ {
		char := words[n][i] - 'a'
		subsetLetters[char]--
		totalScore -= score[char]
	}
}

func isValidWord(freq [26]int, subsetLetters [26]int) bool {
	for c := 0; c < 26; c++ {
		if freq[c] < subsetLetters[c] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 23
	// words := []string{"dog","cat","dad","good"}
	// letters := []byte{'a','a','c','d','d','d','g','o','o'}
	// score := []int{1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0}

	// result: 27
	// words := []string{"xxxz","ax","bx","cx"}
	// letters := []byte{'z','a','b','c','x','x','x'}
	// score := []int{4,4,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,10}

	// result: 0
	words := []string{"leetcode"}
	letters := []byte{'l','e','t','c','o','d'}
	score := []int{0,0,1,1,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,1,0,0,0,0,0,0}

	// result: 
	// words := []string{}
	// letters := []byte{}
	// score := []int{}

	result := maxScoreWords(words, letters, score)
	fmt.Printf("result = %v\n", result)
}

