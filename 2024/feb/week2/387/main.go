package main
import (
	"fmt"
	"strings"
)

const SPACE = byte(rune(' '))

func firstUniqChar(s string) int {
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char == SPACE {
            continue
        }
        if strings.Index(s[i + 1:], string(char)) == -1 {
            return i
        }
        s = strings.Replace(s, string(char), string(SPACE), -1)
    }
    return -1
}

func main() {
	// result: 0
	// s := "leetcode"

	// result: 2
	// s := "loveleetcode"

	// result: -1
	s := "aabb"

	// result: 
	// s := ""

	result := firstUniqChar(s)
	fmt.Printf("result = %v\n", result)
}

