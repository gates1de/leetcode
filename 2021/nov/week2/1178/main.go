package main
import (
	"fmt"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
	result := make([]int, len(puzzles))

	m := make([]map[rune]int, len(words))
	for i, word := range words {
		m[i] = map[rune]int{}
		for _, c := range word {
			m[i][c]++
		}
	}

	for i, puzzle := range puzzles {
		for j, word := range words {
			if isOK(puzzle, word, m[j]) {
				result[i]++
			}
		}
	}

	return result
}

func isOK(puzzle string, word string, m map[rune]int) bool {
	count := int(0)
	for _, char := range puzzle {
		if count == 0 && m[char] == 0 {
			return false
		}
		count += m[char]
		if count == len(word) {
			break
		}
	}
	return count == len(word)
}

func main() {
	// result: [1,1,3,2,4,0]
	// words := []string{"aaaa","asas","able","ability","actt","actor","access"}
	// puzzles := []string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"}

	// result: [0,1,3,2,0]
	words := []string{"apple","pleas","please"}
	puzzles := []string{"aelwxyz","aelpxyz","aelpsxy","saelpxy","xaelpsy"}

	// result: 
	// words := []string{}
	// puzzles := []string{}

	result := findNumOfValidWords(words, puzzles)
	fmt.Printf("result = %v\n", result)
}

