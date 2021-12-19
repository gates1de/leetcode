package main
import (
	"fmt"
	"strconv"
)

const r = byte('[')
const l = byte(']')

func decodeString(s string) string {
	// fmt.Printf("s = %v\n", s)
	result := ""
	repeatCountStr := ""

	for i := 0; i < len(s); i++ {
		c := s[i]
		if byte('0') <= c && c <= byte('9') {
			repeatCountStr += string(c)
		} else if c == r {
			rCount := int(1)
			for j := i + 1; j < len(s); j++ {
				cc := s[j]
				if cc == r {
					rCount++
				} else if cc == l {
					rCount--
				}

				if rCount > 0 {
					continue
				}

				repeatCount, _ := strconv.Atoi(repeatCountStr)
				repeatCountStr = ""
				strInBrackets := decodeString(s[i + 1:j])
				for k := 0; k < repeatCount; k++ {
					result += strInBrackets
				}
				i = j
				break
			}
		} else if byte('a') <= c && c <= byte('z') {
			result += string(c)
		}
	}

	return result
}

func main() {
	// result: "aaabcbc"
	// s := "3[a]2[bc]"

	// result: "accaccacc"
	// s := "3[a2[c]]"

	// result: "abcabccdcdcdef"
	// s := "2[abc]3[cd]ef"

	// result: "abccdcdcdxyz"
	// s := "abc3[cd]xyz"

	// result: "abcdefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmncdefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmncdefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmefghijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklhijkkkkkklmnrstu"
	// s := "ab3[cd4[efg5[hij6[k]l]m]n]rstu"

	// result: "a"
	s := "a"

	// result: 
	// s := ""

	result := decodeString(s)
	fmt.Printf("result = %v\n", result)
}

