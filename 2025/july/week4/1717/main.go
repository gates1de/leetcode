package main

import (
	"fmt"
)

func maximumGain(s string, x int, y int) int {
	if x < y {
		x, y = y, x
		s = reverse(s)
	}

	aCount := int(0)
	bCount := int(0)
	result := int(0)

	for _, char := range s {
		if char == 'a' {
			aCount++
		} else if char == 'b' {
			if aCount > 0 {
				aCount--
				result += x
			} else {
				bCount++
			}
		} else {
			result += min(aCount, bCount) * y
			aCount = 0
			bCount = 0
		}
	}

	result += min(aCount, bCount) * y
	return result
}

func reverse(s string) string {
    chars := []rune(s)
    for i, j := 0, len(chars) - 1; i < j; i, j = i + 1, j - 1 {
        chars[i], chars[j] = chars[j], chars[i]
    }
    return string(chars)
}


func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 19
	// s := "cdbcbbaaabab"
	// x := int(4)
	// y := int(5)

	// result: 20
	s := "aabbaaxybbaabb"
	x := int(5)
	y := int(4)

	// result: 
	// s := ""
	// x := int(0)
	// y := int(0)

	result := maximumGain(s, x, y)
	fmt.Printf("result = %v\n", result)
}
