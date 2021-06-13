package main
import (
	"fmt"
)

// More faster algorithm is Trie Tree
// REF: https://leetcode.com/problems/palindrome-pairs/discuss/158103/Go-trie-O(n*k)-solution-with-explanation

func palindromePairs(words []string) [][]int {
	result := [][]int{}
	for i, w1 := range words {
		for j, w2 := range words {
			if i == j {
				continue
			}
			if isPalindrome(w1 + w2) {
				pair := []int{i, j}
				result = append(result, pair)
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
    if len(s) == 1 {
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

func main() {
	// result: [[0,1],[1,0],[3,2],[2,4]]
	// words := []string{"abcd", "dcba", "lls", "s", "sssll"}

	// result: [[0,1],[1,0]]
	// words := []string{"bat", "tab", "cat"}

	// result: [[0,1],[1,0]]
	words := []string{"a", ""}

	// result: 
	// words := []string{}

	result := palindromePairs(words)
	fmt.Printf("result = %v\n", result)
}

