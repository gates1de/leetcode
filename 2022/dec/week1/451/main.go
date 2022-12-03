package main
import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	m := map[byte]string{}
	for _, char := range s {
		b := byte(char)
		if _, ok := m[b]; !ok {
			m[b] = ""
		}
		m[b] += string(char)
	}

	strs := make([]string, 0, len(m))
	for _, str := range m {
		strs = append(strs, str)
	}

	sort.Slice(strs, func(a, b int) bool {
		return len(strs[a]) > len(strs[b])
	})

	result := ""
	for _, str := range strs {
		result += str
	}
	return result
}

func main() {
	// result: "eert"
	// s := "tree"

	// result: "aaaccc"
	// s := "cccaaa"

	// result: "bbAa"
	s := "Aabb"

	// result: 
	// s := ""

	result := frequencySort(s)
	fmt.Printf("result = %v\n", result)
}

