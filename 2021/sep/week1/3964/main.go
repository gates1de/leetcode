package main
import (
	"fmt"
	"sort"
)

func orderlyQueue(s string, k int) string {
	chars := make([]rune, len(s))
	for i, r := range s {
		chars[i] = r
	}
	sort.Slice(chars, func (a, b int) bool { return chars[a] < chars[b] })
	if len(s) == k || k >= 2 {
		return string(chars)
	}

	index := int(-1)
	result := ""
	for i, r := range s {
		if r != chars[0] {
			continue
		}

		if index == -1 {
			index = i
			result = s[i:] + s[:i]
			continue
		}

		temp := s[i:] + s[:i]
		// fmt.Printf("temp = %v\n", temp)
		if isAscending(temp, result) {
			result = temp
		}
	}
	return result
}

// condition: len(s1) == len(s2)
func isAscending(s1 string, s2 string) bool {
	for i, r1 := range s1 {
		r2 := rune(s2[i])
		if r1 < r2 {
			return true
		} else if r1 > r2 {
			return false
		}
	}
	return true
}

func main() {
	// result: "acb"
	// s := "cba"
	// k := int(1)

	// result: "aaabc"
	// s := "baaca"
	// k := int(3)

	// result: "aaabcdddefffghijnopqs"
	// s := "cbadefhigfjdasoqpfdna"
	// k := int(2)

	// result: "abcbadefhigfjdasoqpfd"
	s := "cbadefhigfjdasoqpfdab"
	k := int(1)

	// result: 
	// s := ""
	// k := int(0)

	result := orderlyQueue(s, k)
	fmt.Printf("result = %v\n", result)
}

