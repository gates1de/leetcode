package main
import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
    if len(s) < len(t) {
		return ""
	}

    m := map[byte]int{}

    for _, v := range []byte(t) {
        m[v]++
    }

    start := int(0)
    end := int(0)
    index := int(0)
    counter := len(t)
    minLen := math.MaxInt32

    for end < len(s) {
        if m[s[end]] > 0 {
            counter--
        }

        m[s[end]]--
        end++

        for counter == 0 {
            if minLen > end - start {
                minLen = end - start
                index = start
            }

            m[s[start]]++
            if m[s[start]] == 1 {
                counter++
            }

            start++
        }
    }

    if minLen == math.MaxInt32 {
        return ""
    }
    return s[index:index + minLen]
}

func main() {
	// result: "BANC"
	s := "ADOBECODEBANC"
	t := "ABC"

	// result: "a"
	// s := "a"
	// t := "a"

	// result: ""
	// s := "a"
	// t := "aa"

	// result: 
	// s := ""
	// t := ""

	result := minWindow(s, t)
	fmt.Printf("result = %v\n", result)
}

