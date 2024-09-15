package main
import (
	"fmt"
)

func findTheLongestSubstring(s string) int {
	prefixXOR := int(0)
	m := make([]int, 26)
	m['a' - 'a'] = 1
	m['e' - 'a'] = 2
	m['i' - 'a'] = 4
	m['o' - 'a'] = 8
	m['u' - 'a'] = 16
	mp := make([]int, 32)

	for i := 0; i < 32; i++ {
		mp[i] = -1
	}

	result := int(0)
	for i := 0; i < len(s); i++ {
		prefixXOR ^= m[s[i] - 'a']

		if mp[prefixXOR] == -1 && prefixXOR != 0 {
			mp[prefixXOR] = i
		}

		result = max(result, i - mp[prefixXOR])
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 13
	// s := "eleetminicoworoep"

	// result: 5
	// s := "eleetminicoworoep"

	// result: 6
	s := "bcbcbc"

	// result: 
	// s := ""

	result := findTheLongestSubstring(s)
	fmt.Printf("result = %v\n", result)
}

