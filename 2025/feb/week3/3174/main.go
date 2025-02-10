package main
import (
	"fmt"
	"strconv"
)

func clearDigits(s string) string {
	resultLen := int(0)
	chars := []byte(s)

	for i, char := range s {
		_, err := strconv.Atoi(string(char))
		if err == nil {
			resultLen = max(resultLen - 1, 0)
		} else {
			chars[resultLen] = s[i]
			resultLen++
		}
	}

	return string(chars[0:resultLen])
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: "abc"
	// s := "abc"

	// result: ""
	s := "cb34"

	// result: ""
	// s := ""

	result := clearDigits(s)
	fmt.Printf("result = %v\n", result)
}

