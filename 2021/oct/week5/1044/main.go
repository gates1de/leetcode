package main
import (
	"fmt"
)

func longestDupSubstring(s string) string {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}

	result := ""
	current := ""
	for i, c := range s {
		if m[c] == 1 {
			current = ""
			continue
		}

		current += string(c)
		targetStr := s[i - len(current) + 2:]
		// fmt.Printf("%v: current = %v, targetStr = %v\n", i, current, targetStr)
		for j := 0; j < len(targetStr) - len(current) + 1; j++ {
			if m[rune(targetStr[j])] == 1 {
				break
			}
			sub := targetStr[j:j + len(current)]
			// fmt.Printf("%v: sub = %v\n", j, sub)
			if current == sub && len(result) < len(current) {
				result = current
			}
		}
	}
	return result
}

func main() {
	// result: "ana"
	// s := "banana"

	// result: ""
	// s := "abcd"

	// result: "fabfa"
	// s := "wiwinfabfabfa"

	// result: "a"
	// s := "aa"

	// result: "ma"
	s := "nnpxouomcofdjuujloanjimymadkuepightrfodmauhrsy"

	// result: 
	// s := ""

	result := longestDupSubstring(s)
	fmt.Printf("result = %v\n", result)
}

