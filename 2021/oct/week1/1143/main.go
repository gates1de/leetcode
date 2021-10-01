package main
import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	long := text1
	short := text2
	if len(text1) < len(text2) {
		long, short = text2, text1
	}

	shortCharIndices := map[rune]int{}
	for i, r := range short {
		shortCharIndices[r] = i
	}

	result := int(0)
	longCharIndices := []int{}
	for _, r := range long {
		if index, ok := shortCharIndices[r]; ok {
			if len(longCharIndices) == 0 || longCharIndices[len(longCharIndices) - 1] < index {
				longCharIndices = append(longCharIndices, index)
				if result < len(longCharIndices) {
					result = len(longCharIndices)
				}
			} else if longCharIndices[len(longCharIndices) - 1] > index {
				for i, val := range longCharIndices {
					if val == index {
						break
					} else if val > index {
						longCharIndices = longCharIndices[:i]
						longCharIndices = append(longCharIndices, index)
						break
					}
				}
			}
			// fmt.Printf("long = %v\n", longCharIndices)
		}
	}

	return result
}

func main() {
	// result: 3
	// text1 := "abcde"
	// text2 := "ace"

	// result: 3
	// text1 := "abc"
	// text2 := "abc"

	// result: 0
	// text1 := "abc"
	// text2 := "def"

	// result: 1
	// text1 := "z"
	// text2 := "z"

	// result: 3
	// text1 := "fhnewalferpwqyiafehsaf"
	// text2 := "fine"

	// result: 6
	// text1 := "mhunuzqrkzsnidwbun"
	// text2 := "szulspmhwpazoxijwbq"

	// result: 5
	text1 := "abcba"
	text2 := "abcbcba"

	// result: 
	// text1 := ""
	// text2 := ""

	result := longestCommonSubsequence(text1, text2)
	fmt.Printf("result = %v\n", result)
}

