package main
import (
	"fmt"
)

func canConstruct(s string, k int) bool {
	n := len(s)

	if n < k {
		return false
	} else if n == k {
		return true
	}

	freq := make([]int, 26)
	oddCount := int(0)

	for _, char := range s {
		freq[char - 'a']++
	}

	for _, count := range freq {
		if count % 2 == 1 {
			oddCount++
		}
	}

	return oddCount <= k
}

func main() {
	// result: true
	// s := "annabelle"
	// k := int(2)

	// result: false
	// s := "leetcode"
	// k := int(3)

	// result: true
	s := "true"
	k := int(4)

	// result: 
	// s := ""
	// k := int(0)

	result := canConstruct(s, k)
	fmt.Printf("result = %v\n", result)
}

