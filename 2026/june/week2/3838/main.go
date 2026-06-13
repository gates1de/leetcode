package main

import (
	"fmt"
)

func mapWordWeights(words []string, weights []int) string {
	result := make([]byte, len(words))
	for i, word := range words {
		sum := int(0)
		for j := range len(word) {
			sum += weights[word[j]-'a']
		}

		result[i] = byte('z' - byte(sum%26))
	}

	return string(result)
}

func main() {
	// result: "rij"
	// words := []string{"abcd", "def", "xyz"}
	// weights := []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}

	// result: "yyy"
	// words := []string{"a", "b", "c"}
	// weights := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	// result: "g"
	words := []string{"abcd"}
	weights := []int{7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5}

	// result: ""
	// words := []string{}
	// weights := []int{}

	result := mapWordWeights(words, weights)
	fmt.Printf("result = %v\n", result)
}
