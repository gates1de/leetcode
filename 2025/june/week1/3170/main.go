package main
import (
	"fmt"
)

func clearStars(s string) string {
	counter := make([][]int, 26)
	for i := range counter {
		counter[i] = make([]int, 0)
	}

	chars := []rune(s)
	for i, c := range chars {
		if c != '*' {
			counter[c - 'a'] = append(counter[c - 'a'], i)
		} else {
			for j := 0; j < 26; j++ {
				if len(counter[j]) > 0 {
					last := len(counter[j]) - 1
					chars[counter[j][last]] = '*'
					counter[j] = counter[j][:last]
					break
				}
			}
		}
	}

	result := make([]rune, 0, len(chars))
	for _, c := range chars {
		if c != '*' {
			result = append(result, c)
		}
	}

	return string(result)
}

func main() {
	// result: "aab"
	// s := "aaba*"

	// result: "abc"
	s := "abc"

	// result: ""
	// s := ""

	result := clearStars(s)
	fmt.Printf("result = %v\n", result)
}

