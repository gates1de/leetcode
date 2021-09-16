package main
import (
	"fmt"
)

func findLUSlength(strs []string) int {
	newStrs := removeDuplicates(strs)
	m := map[int][]string{}
	maxLen := int(0)
	for _, s := range newStrs {
		l := len(s)
		if m[l] == nil {
			m[l] = []string{s}
		} else {
			m[l] = append(m[l], s)
		}
		if maxLen < l {
			maxLen = l
		}
	}

	result := int(-1)
	for i := maxLen; i >= 1; i-- {
		strsPerLen := m[i]
		// fmt.Printf("strsPerLen = %v\n", strsPerLen)
		if strsPerLen == nil {
			continue
		}

		isNotExistsSubseq := true
		MID:
		for _, s1 := range strs {
			for _, s2 := range strsPerLen {
				if len(s1) <= len(s2) || s1 == s2 {
					continue
				}
				// fmt.Printf("s1 = %v, s2 = %v\n", s1, s2)
				if !isSubseq(s1, s2) {
					isNotExistsSubseq = true
					break MID
				}
				isNotExistsSubseq = false
			}
		}
		if isNotExistsSubseq {
			return i
		}
	}

    return result
}

func removeDuplicates(strs []string) []string {
	m := map[string]int{}
	for _, s := range strs {
		m[s]++
	}
	result := []string{}
	for s, count := range m {
		if count != 1 {
			continue
		}
		result = append(result, s)
	}
	return result
}

func isSubseq(w1 string, w2 string) bool {
    i := int(0)
    j := int(0)
    for i < len(w1) && j < len(w2) {
        c1 := w1[i]
        c2 := w2[j]
        if c1 == c2 {
            i++
            j++
        } else {
            i++
        }
    }

    if j == len(w2) {
		return true
    }
    return false
}

func main() {
	// result: 3
	// strs := []string{"aba", "cdc", "eae"}

	// result: -1
	// strs := []string{"aaa", "aaa", "aa"}

	// result: 4
	// strs := []string{"aaa", "aaaa", "aa"}

	// result: -1
	strs := []string{"abcabc", "abcabc", "abcabc", "abc", "abc", "aab"}

	// result: 
	// strs := []string{}

	// result: 
	// strs := []string{}

	result := findLUSlength(strs)
	fmt.Printf("result = %v\n", result)
}

