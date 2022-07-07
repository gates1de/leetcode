package main
import (
	"fmt"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := map[string]bool{}
	return helper(s1, s2, s3, m)
}

func helper(s1 string, s2 string, s3 string, m map[string]bool) bool {
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
        result = result || helper(s1[1:], s2, s3[1:], m)
    }
    if s2[0] == s3[0] {
        result = result || helper(s1, s2[1:], s3[1:], m)
    }
    m[s1 + " " + s2 + " " + s3] = result

    return result
}

// Wrong Answer
func ngSolution(s1 string, s2 string, s3 string) bool {
	s1Index := len(s1) - 1
	s2Index := len(s2) - 1
	for i := len(s3) - 1; i >= 0; i-- {
		s3Char := s3[i]
		fmt.Println(string(s3Char), s1[:s1Index + 1], s2[:s2Index + 1])
		if s1Index >= 0 && s3Char == s1[s1Index] {
			s1Index--
			continue
		}
		if s2Index >= 0 && s3Char == s2[s2Index] {
			s2Index--
			continue
		}
		return false
	}

	return s1Index < 0 && s2Index < 0
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
	// s1 := "a"
	// s2 := "b"
	// s3 := "a"

	// result: true
	s1 := "aabc"
	s2 := "abad"
	s3 := "aabadabc"

	// result: 
	// s1 := ""
	// s2 := ""
	// s3 := ""

	result := isInterleave(s1, s2, s3)
	fmt.Printf("result = %v\n", result)
}

