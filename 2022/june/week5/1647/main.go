package main
import (
	"fmt"
)

func minDeletions(s string) int {
	m := map[rune]int{}
	maxCount := int(0)
	for _, char := range s {
		m[char]++
		if maxCount < m[char] {
			maxCount = m[char]
		}
	}

	result := int(0)
	emptyCounts := []int{}
	for i := 1; i <= maxCount; i++ {
		emptyCounts = append(emptyCounts, i)

		sameCountChars := []rune{}
		for char, count := range m {
			if count != i {
				continue
			}
			sameCountChars = append(sameCountChars, char)
		}

		n := len(sameCountChars)
		if n == 0 {
			continue
		} else if n == 1 {
			emptyCounts = emptyCounts[:len(emptyCounts) - 1]
			continue
		}

		for j, char := range sameCountChars {
			if j < len(emptyCounts) {
				newCount := emptyCounts[len(emptyCounts) - 1 - j]
				result += i - newCount
				m[char] = newCount
			} else {
				result += i
				delete(m, char)
			}
		}
		emptyCounts = emptyCounts[:max(0, len(emptyCounts) - n)]
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
	// result: 0
	// s := "aab"

	// result: 2
	// s := "aaabbbcc"

	// result: 2
	// s := "ceabaacb"

	// result: 15
	// s := "fheowqfenjkfanskfsjkafebjkfaeskafjalsf"

	// result: 1
	s := "accdcdadddbaadbc"

	// result: 
	// s := ""

	result := minDeletions(s)
	fmt.Printf("result = %v\n", result)
}

