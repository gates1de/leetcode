package main

import (
	"fmt"
)

func generateString(str1 string, str2 string) string {
	n := len(str1)
	m := len(str2)
	chars := make([]byte, n + m - 1)
	fixed := make([]int, n + m - 1)

	for i := range chars {
		chars[i] = 'a'
	}

	for i := range n {
		if str1[i] == 'T' {
			for j := i; j < i + m; j++ {
				if fixed[j] == 1 && chars[j] != str2[j - i] {
					return ""
				} else {
					chars[j] = str2[j - i]
					fixed[j] = 1
				}
			}
		}
	}

	for i := range n {
		if str1[i] == 'F' {
			flag := false
			idx := int(-1)
			for j := i + m - 1; j >= i; j-- {
				if str2[j - i] != chars[j] {
					flag = true
				}

				if idx == -1 && fixed[j] == 0 {
					idx = j
				}
			}

			if flag {
				continue
			} else if idx != -1 {
				chars[idx] = 'b'
			} else {
				return ""
			}
		}
	}

	return string(chars)
}

func main() {
	// result: "ababa"
	// str1 := "TFTF"
	// str2 := "ab"

	// result: ""
	// str1 := "TFTF"
	// str2 := "abc"

	// result: "a"
	str1 := "F"
	str2 := "d"

	// result: ""
	// str1 := ""
	// str2 := ""

	result := generateString(str1, str2)
	fmt.Printf("result = %v\n", result)
}

