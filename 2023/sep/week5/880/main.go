package main
import (
	"fmt"
)

func decodeAtIndex(s string, k int) string {
    length := int64(0)
    i := int(0)

    for length < int64(k) {
        if s[i] >= '0' && s[i] <= '9' {
            length *= int64(s[i] - '0')
        } else {
            length++
        }
        i++
    }

    for j := i - 1; j >= 0; j-- {
        if s[j] >= '0' && s[j] <= '9' {
            length /= int64(s[j] - '0')
            k %= int(length)
        } else {
            if k == 0 || k == int(length) {
                return string(s[j])
            }
            length--
        }
    }

    return ""
}

func main() {
	// result: "o"
	// s := "leet2code3"
	// k := int(10)

	// result: "h"
	// s := "ha22"
	// k := int(5)

	// result: "a"
	s := "a2345678999999999999999"
	k := int(1)

	// result: 
	// s := ""
	// k := int(0)

	result := decodeAtIndex(s, k)
	fmt.Printf("result = %v\n", result)
}

