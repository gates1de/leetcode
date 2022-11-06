package main
import (
	"fmt"
	"sort"
)

func orderlyQueue(s string, k int) string {
	chars := []byte(s)
	if k > 1 {
		sort.Slice(chars, func(a, b int) bool { return chars[a] < chars[b] })
		return string(chars)
	}

	n := len(chars)
	resultChars := make([]byte, n)
	copy(resultChars, chars)
	for i := 0; i < n; i++ {
		temp := make([]byte, n)
		copy(temp[:n - i], chars[i:])
		copy(temp[n - i:], chars[:i])
		if compare(temp, resultChars) {
			resultChars = temp
		}
	}

	return string(resultChars)
}

func compare(chars1 []byte, chars2 []byte) bool {
	n := len(chars1)
	if n > len(chars2) {
		n = len(chars2)
	}

	for i := 0; i < n; i++ {
		if chars1[i] < chars2[i] {
			return true
		} else if chars1[i] > chars2[i] {
			return false
		}
	}

	return false
}

func main() {
	// result: "acb"
	// s := "cba"
	// k := int(1)

	// result: "aaabc"
	// s := "baaca"
	// k := int(3)

	// result: "abced"
	s := "dabce"
	k := int(1)

	// result: 
	// s := ""
	// k := int(0)

	result := orderlyQueue(s, k)
	fmt.Printf("result = %v\n", result)
}

