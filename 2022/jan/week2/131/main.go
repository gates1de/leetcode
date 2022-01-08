package main
import (
	"fmt"
)

func partition(s string) [][]string {
	if len(s) <= 1 {
		return [][]string{[]string{s}}
	}

	result := [][]string{}
	helper(s, "", []string{}, &result)
	return result
}

func helper(s string, current string, palindromes []string, result *[][]string) {
	// fmt.Printf("s = %v, current = %v\n", s, current)
	if len(s) == 0 {
		if len(current) > 0 {
			return
		}
		if len(palindromes) > 0 {
			*result = append(*result, palindromes)
		}
		return
	}

	char := string(s[0])
	nextS := ""
	if len(s) > 1 {
		nextS = s[1:]
	}

	if isPalindrome(current + char) {
		newPalindromes := make([]string, len(palindromes) + 1)
		copy(newPalindromes, palindromes)
		newPalindromes[len(newPalindromes) - 1] = current + char
		helper(nextS, "", newPalindromes, result)
	}

	helper(nextS, current + char, palindromes, result)
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	} else if len(s) == 1 {
        return true
    }

    i := int(0)
    j := int(len(s) - 1)
    for i <= j {
        left := s[i]
        right := s[j]
        if left != right {
            return false
        }
        i++
        j--
    }
    return true
}

func sum(strings []string) int {
	result := int(0)
	for _, s := range strings {
		result += len(s)
	}
	return result
}

func main() {
	// result: [["a","a","b"],["aa","b"]]
	// s := "aab"

	// result: [["a"]]
	// s := "a"

	// result: [["a","a","b","b","c","c"],["a","a","b","b","cc"],["a","a","bb","c","c"],["a","a","bb","cc"],["aa","b","b","c","c"],["aa","b","b","cc"],["aa","bb","c","c"],["aa","bb","cc"]]
	// s := "aabbcc"

	// result: [["a","a","a","b","c","a","c","b"],["a","a","a","b","cac","b"],["a","a","a","bcacb"],["a","aa","b","c","a","c","b"],["a","aa","b","cac","b"],["a","aa","bcacb"],["aa","a","b","c","a","c","b"],["aa","a","b","cac","b"],["aa","a","bcacb"],["aaa","b","c","a","c","b"],["aaa","b","cac","b"],["aaa","bcacb"]]
	// s := "aaabcacb"

	// result: [["a","b","a","b","b","b","a","b","b","a","b","a"],["a","b","a","b","b","b","a","b","b","aba"],["a","b","a","b","b","b","a","b","bab","a"],["a","b","a","b","b","b","a","bb","a","b","a"],["a","b","a","b","b","b","a","bb","aba"],["a","b","a","b","b","b","abba","b","a"],["a","b","a","b","b","bab","b","a","b","a"],["a","b","a","b","b","bab","b","aba"],["a","b","a","b","b","bab","bab","a"],["a","b","a","b","b","babbab","a"],["a","b","a","b","bb","a","b","b","a","b","a"],["a","b","a","b","bb","a","b","b","aba"],["a","b","a","b","bb","a","b","bab","a"],["a","b","a","b","bb","a","bb","a","b","a"],["a","b","a","b","bb","a","bb","aba"],["a","b","a","b","bb","abba","b","a"],["a","b","a","b","bbabb","a","b","a"],["a","b","a","b","bbabb","aba"],["a","b","a","bb","b","a","b","b","a","b","a"],["a","b","a","bb","b","a","b","b","aba"],["a","b","a","bb","b","a","b","bab","a"],["a","b","a","bb","b","a","bb","a","b","a"],["a","b","a","bb","b","a","bb","aba"],["a","b","a","bb","b","abba","b","a"],["a","b","a","bb","bab","b","a","b","a"],["a","b","a","bb","bab","b","aba"],["a","b","a","bb","bab","bab","a"],["a","b","a","bb","babbab","a"],["a","b","a","bbb","a","b","b","a","b","a"],["a","b","a","bbb","a","b","b","aba"],["a","b","a","bbb","a","b","bab","a"],["a","b","a","bbb","a","bb","a","b","a"],["a","b","a","bbb","a","bb","aba"],["a","b","a","bbb","abba","b","a"],["a","b","abbba","b","b","a","b","a"],["a","b","abbba","b","b","aba"],["a","b","abbba","b","bab","a"],["a","b","abbba","bb","a","b","a"],["a","b","abbba","bb","aba"],["a","bab","b","b","a","b","b","a","b","a"],["a","bab","b","b","a","b","b","aba"],["a","bab","b","b","a","b","bab","a"],["a","bab","b","b","a","bb","a","b","a"],["a","bab","b","b","a","bb","aba"],["a","bab","b","b","abba","b","a"],["a","bab","b","bab","b","a","b","a"],["a","bab","b","bab","b","aba"],["a","bab","b","bab","bab","a"],["a","bab","b","babbab","a"],["a","bab","bb","a","b","b","a","b","a"],["a","bab","bb","a","b","b","aba"],["a","bab","bb","a","b","bab","a"],["a","bab","bb","a","bb","a","b","a"],["a","bab","bb","a","bb","aba"],["a","bab","bb","abba","b","a"],["a","bab","bbabb","a","b","a"],["a","bab","bbabb","aba"],["a","babbbab","b","a","b","a"],["a","babbbab","b","aba"],["a","babbbab","bab","a"],["aba","b","b","b","a","b","b","a","b","a"],["aba","b","b","b","a","b","b","aba"],["aba","b","b","b","a","b","bab","a"],["aba","b","b","b","a","bb","a","b","a"],["aba","b","b","b","a","bb","aba"],["aba","b","b","b","abba","b","a"],["aba","b","b","bab","b","a","b","a"],["aba","b","b","bab","b","aba"],["aba","b","b","bab","bab","a"],["aba","b","b","babbab","a"],["aba","b","bb","a","b","b","a","b","a"],["aba","b","bb","a","b","b","aba"],["aba","b","bb","a","b","bab","a"],["aba","b","bb","a","bb","a","b","a"],["aba","b","bb","a","bb","aba"],["aba","b","bb","abba","b","a"],["aba","b","bbabb","a","b","a"],["aba","b","bbabb","aba"],["aba","bb","b","a","b","b","a","b","a"],["aba","bb","b","a","b","b","aba"],["aba","bb","b","a","b","bab","a"],["aba","bb","b","a","bb","a","b","a"],["aba","bb","b","a","bb","aba"],["aba","bb","b","abba","b","a"],["aba","bb","bab","b","a","b","a"],["aba","bb","bab","b","aba"],["aba","bb","bab","bab","a"],["aba","bb","babbab","a"],["aba","bbb","a","b","b","a","b","a"],["aba","bbb","a","b","b","aba"],["aba","bbb","a","b","bab","a"],["aba","bbb","a","bb","a","b","a"],["aba","bbb","a","bb","aba"],["aba","bbb","abba","b","a"]]
	s := "ababbbabbaba"

	// result: 
	// s := ""

	result := partition(s)
	fmt.Printf("result = %v\n", result)
}

