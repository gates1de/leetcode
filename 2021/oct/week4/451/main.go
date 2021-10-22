package main
import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	weights := map[rune]int{}
	for _, c := range s {
		weights[c]++
	}

	chars := []rune(s)
	sort.Slice(chars, func (a, b int) bool {
		if weights[chars[a]] == weights[chars[b]] {
			return chars[a] < chars[b]
		}
		return weights[chars[a]] > weights[chars[b]]
	})

	return string(chars)
}

func main() {
	// result: "eert"
	// s := "tree"

	// result: "aaaccc"
	// s := "cccaaa"

	// result: "bbAa"
	// s := "Aabb"

	// result: "aaaaaaddddddssssssbbbbbfffffkkkkkjjjjllen"
	s := "safselkjafbndsjkabfkjdsbafdbkjadslbdkafds"

	// result: 
	// s := ""

	result := frequencySort(s)
	fmt.Printf("result = %v\n", result)
}

