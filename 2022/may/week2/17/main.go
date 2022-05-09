package main
import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	first := digits[0]
	strings := getLetters(first)
	return helper(digits[1:], strings)
}

func helper(digits string, strings []string) []string {
	if len(digits) == 0 {
		return strings
	}

	digit := digits[0]
	letters := getLetters(digit)
	newStrings := make([]string, len(strings) * len(letters))
	for i, str := range strings {
		for j, letter := range letters {
			newStrings[i * len(letters) + j] = str + letter
		}
	}

	return helper(digits[1:], newStrings)
}

func getLetters(b byte) []string {
	letters := map[byte][]string{
		byte('2'): {"a", "b", "c"},
		byte('3'): {"d", "e", "f"},
		byte('4'): {"g", "h", "i"},
		byte('5'): {"j", "k", "l"},
		byte('6'): {"m", "n", "o"},
		byte('7'): {"p", "q", "r", "s"},
		byte('8'): {"t", "u", "v"},
		byte('9'): {"w", "x", "y", "z"},
	}
	return letters[b]
}

func main() {
	// result: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	// digits := "23"

	// result: []
	// digits := ""

	// result: ["a","b","c"]
	// digits := "2"

	// result: ["adg","adh","adi","aeg","aeh","aei","afg","afh","afi","bdg","bdh","bdi","beg","beh","bei","bfg","bfh","bfi","cdg","cdh","cdi","ceg","ceh","cei","cfg","cfh","cfi"]
	digits := "234"

	// result: 
	// digits := ""

	result := letterCombinations(digits)
	fmt.Printf("result = %v\n", result)
}

