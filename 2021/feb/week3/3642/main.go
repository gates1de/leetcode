package main
import (
	"fmt"
)

func letterCasePermutation(S string) []string {
	result := []string{}
	permute(S, 0, len(S), &result)
	return result
}

func permute(S string, i int, len int, result *[]string) {
	if i == len {
		*result = append(*result, S)
		return
	}

	upperA := rune('A')
	upperZ := rune('Z')
	lowerA := rune('a')
	lowerZ := rune('z')
	runeS := []rune(S)
	r := runeS[i]
	if upperA <= r && r <= upperZ {
		permute(string(runeS), i + 1, len, result)
		runeS[i] = r + 32
		permute(string(runeS), i + 1, len, result)
	} else if lowerA <= r && r <= lowerZ {
		permute(string(runeS), i + 1, len, result)
		runeS[i] = r - 32
		permute(string(runeS), i + 1, len, result)
	} else {
		permute(S, i + 1, len, result)
	}
}

func main() {
	// S := "a1b2"
	// S := "3z4"
	// S := "12345"
	S := "0"

	result := letterCasePermutation(S)
	fmt.Printf("result = %v\n", result)
}

