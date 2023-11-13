package main
import (
	"fmt"
	"sort"
)

func sortVowels(s string) string {
	vowelsAndConsonantsMap := map[byte]bool{
		65: true,
		69: true,
		73: true,
		79: true,
		85: true,
		97: true,
		101: true,
		105: true,
		111: true,
		117: true,
	}

	chars := []byte(s)
	vowelsAndConsonants := make([]byte, 0, len(s))
	for _, char := range chars {
		if vowelsAndConsonantsMap[char] {
			vowelsAndConsonants = append(vowelsAndConsonants, char)
		}
	}

	if len(vowelsAndConsonants) == 0 {
		return s
	}

	sort.Slice(vowelsAndConsonants, func (a, b int) bool {
		return vowelsAndConsonants[a] < vowelsAndConsonants[b]
	})

	replaceIndex := 0
	for i, char := range chars {
		if vowelsAndConsonantsMap[char] {
			chars[i] = vowelsAndConsonants[replaceIndex]
			replaceIndex++
		}
	}

	return string(chars)
}

func main() {
	// result: "lEOtcede"
	// s := "lEetcOde"

	// result: "lYmpH"
	s := "lYmpH"

	// result: ""
	// s := ""

	result := sortVowels(s)
	fmt.Printf("result = %v\n", result)
}

