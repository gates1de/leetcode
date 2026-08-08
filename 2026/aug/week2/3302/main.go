package main

import (
	"fmt"
)

func validSequence(word1 string, word2 string) []int {
	n, m := len(word1), len(word2)

	suffixStart := make([]int, m+1)
	suffixStart[m] = n
	i, j := n-1, m-1
	for j >= 0 {
		for i >= 0 && word1[i] != word2[j] {
			i--
		}

		if i < 0 {
			break
		}

		suffixStart[j] = i
		i--
		j--
	}

	for ; j >= 0; j-- {
		suffixStart[j] = n
	}

	result := make([]int, 0, m)
	usedChange := false
	for i, c := range []byte(word1) {
		p := len(result)
		if p == m {
			break
		}

		if c == word2[p] {
			result = append(result, i)
			continue
		}

		if !usedChange && (p+1 == m || (suffixStart[p+1] < n && suffixStart[p+1] > i)) {
			result = append(result, i)
			usedChange = true
		}
	}

	if len(result) != m {
		return []int{}
	}
	return result
}

func main() {
	// result: [0,1,2]
	// word1 := "vbcca"
	// word2 := "abc"

	// result: [1,2,4]
	// word1 := "bacdc"
	// word2 := "abc"

	// result: []
	// word1 := "aaaaaa"
	// word2 := "aaabc"

	// result: [0,1]
	word1 := "abc"
	word2 := "ab"

	// result: []
	// word1 := ""
	// word2 := ""

	result := validSequence(word1, word2)
	fmt.Printf("result = %v\n", result)
}
