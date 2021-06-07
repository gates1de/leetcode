package main
import (
	"fmt"
)

var m map[string]bool = map[string]bool{}

func isInterleave(s1, s2, s3 string) bool {
	// fmt.Printf("m = %v\n", m)
	if len(s3) != len(s1) + len(s2) {
		return false
	}
    if len(s1) == 0 {
        return s2 == s3
    }
    if len(s2) == 0 {
        return s1 == s3
    }
    if prev, ok := m[s1 + " " + s2 + " " + s3]; ok {
        return prev
    }
    result := false
    if s1[0] == s3[0] {
        result = result || isInterleave(s1[1:], s2, s3[1:])
    }
    if s2[0] == s3[0] {
        result = result || isInterleave(s1, s2[1:], s3[1:])
    }
    m[s1 + " " + s2 + " " + s3] = result
    return result
}

// Slow & Use more memory
func mySolution(s1 string, s2 string, s3 string) bool {
	// fmt.Printf("s1 = %v, s2 = %v, s3 = %v\n", s1, s2, s3)
	if len(s3) != len(s1) + len(s2) {
		return false
	}
	if len(s3) == 0 {
		return true
	}

	result := false
	r := rune(s3[0])
	if len(s1) > 0 && rune(s1[0]) == r {
		result = isInterleave(s1[1:], s2, s3[1:])
	}
	if result {
		return true
	}
	if len(s2) > 0 && rune(s2[0]) == r {
		return isInterleave(s1, s2[1:], s3[1:])
	}
	return false
}

func main() {
	// result: true
	// s1 := "aabcc"
	// s2 := "dbbca"
	// s3 := "aadbbcbcac"

	// result: false
	// s1 := "aabcc"
	// s2 := "dbbca"
	// s3 := "aadbbbaccc"

	// result: true
	// s1 := ""
	// s2 := ""
	// s3 := ""

	// result: false
	s1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s3 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	// result: 
	// s1 := ""
	// s2 := ""
	// s3 := ""

	result := isInterleave(s1, s2, s3)
	fmt.Printf("result = %v\n", result)
}

