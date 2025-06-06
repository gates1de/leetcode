package main
import (
	"fmt"
)

func robotWithString(s string) string {
	counter := make([]int, 26)
	for _, c := range s {
		counter[c - 'a']++
	}

	stack := make([]rune, 0, len(s))
	result := make([]rune, 0, len(s))
	minChar := 'a'

	for _, c := range s {
		stack = append(stack, c)
		counter[c-'a']--
		for minChar != 'z' && counter[minChar - 'a'] == 0 {
			minChar++
		}

		for len(stack) > 0 && stack[len(stack) - 1] <= minChar {
			result = append(result, stack[len(stack) - 1])
			stack = stack[:len(stack) - 1]
		}
	}

	return string(result)
}

func main() {
	// result: "azz"
	// s := "zza"

	// result: "abc"
	// s := "bac"

	// result: "addb"
	s := "bdda"

	// result: ""
	// s := ""

	result := robotWithString(s)
	fmt.Printf("result = %v\n", result)
}

