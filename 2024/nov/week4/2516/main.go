package main
import (
	"fmt"
)

func takeCharacters(s string, k int) int {
	count := make([]int, 3)
	n := len(s)

	for _, char := range s {
		count[char - 'a']++
	}

	for _, c := range count {
		if c < k {
			return -1
		}
	}

	window := make([]int, 3)
	left := int(0)
	maxWindow := int(0)

	for right := 0; right < n; right++ {
		window[s[right] - 'a']++

		for left <= right &&
			(count[0] - window[0] < k || count[1] - window[1] < k || count[2] - window[2] < k) {
				window[s[left] - 'a']--
				left++
		}

		maxWindow = max(maxWindow, right - left + 1)
	}

	return n - maxWindow
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 8
	// s := "aabaaaacaabc"
	// k := int(2)

	// result: -1
	s := "a"
	k := int(1)

	// result: 
	// s := ""
	// k := int(0)

	result := takeCharacters(s, k)
	fmt.Printf("result = %v\n", result)
}

