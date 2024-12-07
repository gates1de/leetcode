package main
import (
	"fmt"
)

func addSpaces(s string, spaces []int) string {
	chars := make([]byte, 0, len(s) + len(spaces))
	for i, char := range s {
		if len(spaces) > 0 && i == spaces[0] {
			chars = append(chars, ' ')
			spaces = spaces[1:]
		}

		chars = append(chars, byte(char))
	}

	return string(chars)
}

func main() {
	// result: "Leetcode Helps Me Learn"
	// s := "LeetcodeHelpsMeLearn"
	// spaces := []int{8,13,15}

	// result: "i code in py thon"
	// s := "icodeinpython"
	// spaces := []int{1,5,7,9}

	// result: " s p a c i n g"
	s := "spacing"
	spaces := []int{0,1,2,3,4,5,6}

	// result: ""
	// s := ""
	// spaces := []int{}

	result := addSpaces(s, spaces)
	fmt.Printf("result = %v\n", result)
}

