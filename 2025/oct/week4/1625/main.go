package main

import (
	"fmt"
)

func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	visited := make([]bool, n)
	result := s

	s = s + s
	for i := 0; !visited[i]; i = (i + b) % n {
		visited[i] = true

		for j := 0; j < 10; j++ {
			limit := 0
			if b % 2 != 0 {
				limit = 9
			}

			for k := 0; k <= limit; k++ {
				t := []byte(s[i : i+n])

				for p := 1; p < n; p += 2 {
					t[p] = '0' + byte((int(t[p] - '0')+ j * a) % 10)
				}

				for p := 0; p < n; p += 2 {
					t[p] = '0' + byte((int(t[p] - '0') + k * a) % 10)
				}

				str := string(t)
				if str < result {
					result = str
				}
			}
		}
	}

	return result
}

func main() {
	// result: "2050"
	// s := "5525"
	// a := int(9)
	// b := int(2)

	// result: "24"
	// s := "74"
	// a := int(5)
	// b := int(1)

	// result: "0011"
	s := "0011"
	a := int(4)
	b := int(2)

	// result: ""
	// s := ""
	// a := int(0)
	// b := int(0)

	result := findLexSmallestString(s, a, b)
	fmt.Printf("result = %v\n", result)
}

