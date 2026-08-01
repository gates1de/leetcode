package main

import (
	"fmt"
)

const alphabetSize = 26

func countPermutations(counts [alphabetSize]int, limit int64) int64 {
	result := int64(1)
	total := int(0)

	for _, count := range counts {
		if count == 0 {
			continue
		}

		choose := min(total, count)
		combination := int64(1)
		for index := range choose {
			numerator := int64(total + count - choose + index + 1)
			denominator := int64(index + 1)
			combination = combination * numerator / denominator
			if combination > limit {
				return limit + 1
			}
		}

		if result > limit / combination {
			return limit + 1
		}
		result *= combination
		total += count
	}

	return result
}

func smallestPalindrome(s string, k int) string {
	counts := [alphabetSize]int{}
	for _, character := range s {
		counts[character - 'a']++
	}

	halfCounts := counts
	for character, count := range counts {
		halfCounts[character] = count / 2
	}

	if countPermutations(halfCounts, int64(k)) < int64(k) {
		return ""
	}

	half := make([]byte, 0, len(s) / 2)
	for range len(s) / 2 {
		for character, count := range halfCounts {
			if count == 0 {
				continue
			}

			halfCounts[character]--
			result := countPermutations(halfCounts, int64(k))
			if result >= int64(k) {
				half = append(half, byte('a'+character))
				break
			}

			k -= int(result)
			halfCounts[character]++
		}
	}

	resultBytes := make([]byte, len(s))
	copy(resultBytes, half)
	for character, count := range counts {
		if count % 2 == 1 {
			resultBytes[len(s) / 2] = byte('a' + character)
			break
		}
	}

	for index := range half {
		resultBytes[len(s) - 1 - index] = half[index]
	}

	return string(resultBytes)
}

func main() {
	// result: "baab"
	// s := "abba"
	// k := int(2)

	// result: ""
	// s := "aa"
	// k := int(2)

	// result: "abcba"
	s := "bacab"
	k := int(1)

	// result: ""
	// s := ""
	// k := int(0)

	result := smallestPalindrome(s, k)
	fmt.Printf("result = %v\n", result)
}
