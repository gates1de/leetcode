package main
import (
	"fmt"
)

func shortestToChar(s string, c byte) []int {
	result := make([]int, len(s))
	leftCharIndex := int(-1)

	for i, char := range s {
		leftDist := int(0)
		rightDist := int(0)

		if c == byte(char) {
			leftCharIndex = i
		}

		if leftCharIndex >= 0 {
			leftDist = i - leftCharIndex
		}
		isExistsOnRight := false
		for j, rightChar := range s[i + 1:] {
			if c == byte(rightChar) {
				rightDist = j + 1
				isExistsOnRight = true
				break
			}
		}

		if !isExistsOnRight {
			result[i] = leftDist
			continue
		}

		if leftCharIndex < 0 {
			result[i] = rightDist
		} else {
			result[i] = min(leftDist, rightDist)
		}
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// s := "loveleetcode"
	// c := []byte("e")

	// s := "aaab"
	// c := []byte("b")

	// s := "aaba"
	// c := []byte("b")

	s := "abaa"
	c := []byte("b")

	result := shortestToChar(s, c[0])
	fmt.Printf("result = %v\n", result)
}

