package main
import (
	"fmt"
)

func largestVariance(s string) int {
	counter := make([]int, 26)
	for _, char := range s {
		counter[char - 'a']++
	}

	result := int(0)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i == j || counter[i] == 0 || counter[j] == 0 {
				continue
			}

			major := byte('a' + i)
			minor := byte('a' + j)
			majorCount := int(0)
			minorCount := int(0)
			restMinor := counter[j]

			for _, char := range s {
				if byte(char) == major {
					majorCount++
				}
				if byte(char) == minor {
					minorCount++
					restMinor--
				}

				if minorCount > 0 {
					result = max(result, majorCount - minorCount)
				}
				if majorCount < minorCount && restMinor > 0 {
					majorCount = 0
					minorCount = 0
				}
			}
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
	// result: 3
	// s := "aababbb"

	// result: 0
	s := "abcde"

	// result: 
	// s := ""

	result := largestVariance(s)
	fmt.Printf("result = %v\n", result)
}

