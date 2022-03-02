package main
import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	} else if len(s) == 0 {
		return true
	} else if len(t) == 0 {
		return false
	} else if len(s) == len(t) {
		return s == t
	}

	i := int(0)
	for _, tChar := range t {
		sChar := s[i]
		if rune(sChar) == tChar {
			i++
		}
		if i == len(s) {
			return true
		}
	}

	return false
}

func main() {
	// result: true
	// s := "abc"
	// t := "ahbgdc"

	// result: false
	// s := "axc"
	// t := "ahbgdc"

	// result: false
	// s := "albmcn"
	// t := "abcdefghijklmnopqrstuvwxyz"

	// result: false
	// s := "abc"
	// t := ""

	// result: true
	// s := ""
	// t := "abc"

	// result: true
	// s := ""
	// t := ""

	// result: false
	s := "acb"
	t := "ahbgdc"

	// result: 
	// s := ""
	// t := ""

	result := isSubsequence(s, t)
	fmt.Printf("result = %v\n", result)
}

