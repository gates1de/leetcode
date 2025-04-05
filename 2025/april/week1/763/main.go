package main
import (
	"fmt"
)

func partitionLabels(s string) []int {
	lastOccurrence := make([]int, 26)
	for i, char := range s {
		lastOccurrence[char - 'a'] = i
	}

	partitionStart := int(0)
	partitionEnd := int(0)
	result := make([]int, 0)
	for i, char := range s {
		partitionEnd = max(partitionEnd, lastOccurrence[char - 'a'])

		if i == partitionEnd {
			result = append(result, i - partitionStart + 1)
			partitionStart = i + 1
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result:[9.7,8] 
	// s := "ababcbacadefegdehijhklij"

	// result: [10]
	s := "eccbbbbdec"

	// result: []
	// s := ""

	result := partitionLabels(s)
	fmt.Printf("result = %v\n", result)
}

