package main
import (
	"fmt"
)

func minimumLength(s string) int {
	return deleteSimilarEnds(s, 0, len(s) - 1)
}

func deleteSimilarEnds(s string, begin int, end int) int {
	if begin >= end || s[begin] != s[end] {
		return end - begin + 1
	}

	char := s[begin]
	for begin <= end && s[begin] == char {
		begin++
	}
	for end > begin && s[end] == char {
		end--
	}

	return deleteSimilarEnds(s, begin, end)
}

func main() {
	// result: 2
	// s := "ca"

	// result: 0
	// s := "cabaabac"

	// result: 3
	s := "aabccabba"

	// result: 
	// s := ""

	result := minimumLength(s)
	fmt.Printf("result = %v\n", result)
}

