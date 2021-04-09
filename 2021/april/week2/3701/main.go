package main
import (
	"fmt"
)

var M = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	} else if len(digits) == 1 {
		return M[digits]
	}

	result := []string{}
	concat(digits, "", &result, len(digits))
	return result
}

func concat(digits string, current string, result *[]string, digitsLen int) {
	// fmt.Printf("current = %v, digits = %v\n", current, digits)
	if len(digits) == 0 {
		if len(current) == digitsLen {
			*result = append(*result, current)
		}
		return
	}

	for i, digit := range digits {
		alphabets := M[string(digit)]
		for _, alphabet := range alphabets {
			concat(digits[i + 1:], current + alphabet, result, digitsLen)
		}
	}
}

func main() {
	// result: [ad, ae, af, bd, be, bf, cd, ce, cf]
	// digits := "23"

	// result: []
	// digits := ""

	// result: [a, b, c]
	// digits := "2"

	// result: ["adg","adh","adi","aeg","aeh","aei","afg","afh","afi","bdg","bdh","bdi","beg","beh","bei","bfg","bfh","bfi","cdg","cdh","cdi","ceg","ceh","cei","cfg","cfh","cfi"]
	digits := "234"

	result := letterCombinations(digits)
	fmt.Printf("result = %v\n", result)
}

