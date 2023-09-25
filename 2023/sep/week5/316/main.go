package main
import (
	"fmt"
)

func removeDuplicateLetters(s string) string {
    m := map[byte]int{}
    n := len(s)
    for i := 0; i < n; i++ {
        m[s[i]]++
    }

    set := map[byte]bool{}

    var stack []byte
    for i := 0; i < n; i++ {
        m[s[i]]--
        stackLen := len(stack)
        if !set[s[i]] {
            for stackLen > 0 && stack[stackLen - 1] > s[i] && m[stack[stackLen - 1]] > 0 {
                set[stack[stackLen - 1]] = false
                stack = stack[:stackLen - 1]
                stackLen = len(stack)
            }
            stack = append(stack, s[i])
            set[s[i]] = true
        }
    }
    return string(stack)
}

func main() {
	// result: "abc"
	// s := "bcabc"

	// result: "acdb"
	s := "cbacdcbc"

	// result: 
	// s := ""

	result := removeDuplicateLetters(s)
	fmt.Printf("result = %v\n", result)
}

