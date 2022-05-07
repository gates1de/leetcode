package main
import (
	"fmt"
)

func removeDuplicates(s string, k int) string {
	if len(s) <= 1 {
		return s
	}

	pre := s[0]
	count := int(1)
	for i := 1; i < len(s); i++ {
		fmt.Println(i, s)
		current := s[i]
		if pre == current {
			count++
		} else {
			pre = current
			count = 1
			continue
		}

		if count == k {
			s = string(s[:i - k + 1]) + string(s[i + 1:])

			if len(s) == 0 {
				break
			}

			i -= 2 * k
			if i < 0 {
				i = 0
			}
			pre = s[i]
			count = 1
		}
	}

	return s
}

func main() {
	// result: "abcd"
	// s := "abcd"
	// k := int(2)

	// result: "aa"
	// s := "deeedbbcccbdaa"
	// k := int(3)

	// result: "ps"
	s := "pbbcggttciiippooaais"
	k := int(2)

	// result: 
	// s := ""
	// k := int(0)

	result := removeDuplicates(s, k)
	fmt.Printf("result = %v\n", result)
}

