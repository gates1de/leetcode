package main
import (
	"fmt"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    str1 := ""
    str2 := ""

    n := len(word1)
    m := len(word2)
    limit := n
    if n < m {
        limit = m
    }

    i := int(0)
    j := int(0)
    for i < limit && j < limit {
        if i < n {
            str1 += word1[i]
            i++
        }
        if j < m {
            str2 += word2[j]
            j++
        }
    }

    return str1 == str2
}

func main() {
	// result: true
	// word1 := []string{"ab", "c"}
	// word2 := []string{"a", "bc"}

	// result: false
	// word1 := []string{"a", "cb"}
	// word2 := []string{"ab", "c"}

	// result: true
	word1 := []string{"abc", "d", "defg"}
	word2 := []string{"abcddefg"}

	// result: 
	// word1 := []string{}
	// word2 := []string{}

	result := arrayStringsAreEqual(word1, word2)
	fmt.Printf("result = %v\n", result)
}

