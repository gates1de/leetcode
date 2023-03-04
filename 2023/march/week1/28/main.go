package main
import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)
	if n1 == 0 || n2 == 0 || n1 < n2 {
		return -1
	}

	TOP:
	for i, char1 := range haystack {
		char2 := rune(needle[0])
		if char1 != char2 {
			continue
		}

		for j := 0; j < n2; j++ {
			temp := haystack[i:]
			if len(temp) < n2 {
				break TOP
			}

			char1 = rune(temp[j])
			char2 = rune(needle[j])
			if char1 != char2 {
				continue TOP
			}
		}
		return i
	}

	return -1
}

func main() {
	// result: 0
	// haystack := "sadbutsad"
	// needle := "sad"

	// result: -1
	// haystack := "leetcode"
	// needle := "leeto"

	// result: 
	haystack := "mississippi"
	needle := "issipi"

	// result: 
	// haystack := ""
	// needle := ""

	result := strStr(haystack, needle)
	fmt.Printf("result = %v\n", result)
}

