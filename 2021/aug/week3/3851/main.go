package main
import (
	"fmt"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	} else if len(s) == len(t) {
		if isIncluded(s, t) {
			return s
		}
		return ""
	}

	result := ""
	l := int(0)
	r := len(t)
	for l <= len(s) - len(t) {
		S := s[l:r]
		// fmt.Printf("l: %v, r: %v, S = %v\n", l, r, S)
		if isIncluded(t, S) {
			if len(S) == len(t) {
				return S
			}

			if result == "" {
				result = S
			} else {
				if len(result) > len(S) {
					result = S
				}
			}
			l++
		} else {
			if r - l == len(s) {
				return ""
			}

			if r < len(s) {
				r++
			} else {
				l++
			}
		}
	}

	return result
}

func isIncluded(target string, s string) bool {
    dict := make(map[rune]int)
    for _, v := range target {
        dict[v]++
    }
    for _, v := range s {
        dict[v]--
    }
    for k, _ := range dict {
        if dict[k] > 0 {
            return false
        }
    }
    return true
}

func main() {
	// result: "BANC"
	// s := "ADOBECODEBANC"
	// t := "ABC"

	// result: "a"
	// s := "a"
	// t := "a"

	// result: ""
	// s := "a"
	// t := "aa"

	// result: "b"
	s := "ab"
	t := "b"

	// result: 
	// s := ""
	// t := ""

	// result: 
	// s := ""
	// t := ""

	result := minWindow(s, t)
	fmt.Printf("result = %v\n", result)
}

