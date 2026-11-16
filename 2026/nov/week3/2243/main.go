package main

import (
	"fmt"
	"strconv"
)

func digitSum(s string, k int) string {
	for len(s) > k {
		next := make([]byte, 0, (len(s)+k-1)/k*3)
		for start := 0; start < len(s); start += k {
			end := min(start+k, len(s))
			sum := 0
			for _, digit := range s[start:end] {
				sum += int(digit - '0')
			}
			next = strconv.AppendInt(next, int64(sum), 10)
		}
		s = string(next)
	}

	return s
}

func main() {
	// result: "135"
	// s := "11111222223"
	// k := 3

	// result: "000"
	s := "00000000"
	k := 3

	// result: ""
	// s := ""
	// k := 0

	result := digitSum(s, k)
	fmt.Printf("result = %v\n", result)
}
