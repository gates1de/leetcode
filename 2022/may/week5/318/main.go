package main
import (
	"fmt"
	"sort"
)

func maxProduct(words []string) int {
    sort.Slice(words, func(a, b int) bool {
        return len(words[a]) > len(words[b])
    })

    alpha := make([]int, len(words))
    for i, w := range words {
        for _, c := range w {
            alpha[i] = alpha[i] | (0x1 << uint(c - 'a'))
        }
    }

	result := int(0)
    for i := 0; i < len(words); i++ {
        if result >= len(words[i]) * len(words[i]) {
            break
        }

        for j := i + 1; j < len(words); j++ {
            if result >= len(words[i]) * len(words[j]) {
                break
            }

            if alpha[i] & alpha[j] != 0 {
                continue
            }

            result = len(words[i]) * len(words[j])
            break
        }
    }

    return result
}

func main() {
	// result: 16
	// words := []string{"abcw","baz","foo","bar","xtfn","abcdef"}

	// result: 4
	// words := []string{"a","ab","abc","d","cd","bcd","abcd"}

	// result: 0
	// words := []string{"a","aa","aaa","aaaa"}

	// result: 15
	words := []string{"eae","ea","aaf","bda","fcf","dc","ac","ce","cefde","dabae"}

	// result: 
	// words := []string{}

	result := maxProduct(words)
	fmt.Printf("result = %v\n", result)
}

