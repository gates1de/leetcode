package main
import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
    s += "0"
    i := int(0)
    j := int(0)
    m := make(map[byte]int)
    for i < len(s) {
        sChar := s[i]
        tChar := t[j]

        m[sChar]++
        m[tChar]--

        i++
        j++
    }

    a := byte('a')
    for k := byte(0); k < 26; k++ {
        if m[a + k] != 0 {
            return a + k
        }
    }

    return byte(0)
}

func main() {
	// result: "e"
	// s := "abcd"
	// t := "abcde"

	// result: "y"
	s := ""
	t := "y"

	// result: 
	// s := ""
	// t := ""

	result := findTheDifference(s, t)
	fmt.Printf("result = %v\n", result)
}

