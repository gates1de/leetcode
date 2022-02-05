package main
import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	counter := map[rune]int{}
	for _, char := range p {
		counter[char]--
	}

	result := []int{}
	for _, char := range s[:len(p)] {
		counter[char]++
		if counter[char] == 0 {
			delete(counter, char)
		}
	}

	if len(counter) == 0 {
		result = append(result, 0)
	}

	for i := 1; i <= len(s) - len(p); i++ {
		out := rune(s[i - 1])
		in := rune(s[i + len(p) - 1])
		// fmt.Println(out, string(out), in, string(in))

		counter[out]--
		counter[in]++
		if counter[out] == 0 {
			delete(counter, out)
		}
		if counter[in] == 0 {
			delete(counter, in)
		}

		// fmt.Println(counter)
		if len(counter) == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	// result: [0,6]
	// s := "cbaebabacd"
	// p := "abc"

	// result: [0,1,2]
	// s := "abab"
	// p := "ab"

	// result: [0]
	// s := "a"
	// p := "a"

	// result: []
	// s := "a"
	// p := "b"

	// result: [0,2,3]
	// s := "acebdfhkigjcae"
	// p := "abcdefghijk"

	// result: [1]
	// s := "baa"
	// p := "aa"

	// result: []
	s := "aaaaaaaaaa"
	p := "aaaaaaaaaaaaa"

	// result: 
	// s := ""
	// p := ""

	result := findAnagrams(s, p)
	fmt.Printf("result = %v\n", result)
}

