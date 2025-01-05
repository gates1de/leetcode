package main
import (
	"fmt"
)

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diffs := make([]int, n)

	for _, shift := range shifts {
		if shift[2] == 1 {
			diffs[shift[0]]++

			if shift[1] + 1 < n {
				diffs[shift[1] + 1]--
			}
		} else {
			diffs[shift[0]]--

			if shift[1] + 1 < n {
				diffs[shift[1] + 1]++
			}
		}
	}

	result := []byte(s)
	numOfShifts := int(0)

	for i := 0; i < n; i++ {
		numOfShifts = (numOfShifts + diffs[i]) % 26

		if numOfShifts < 0 {
			numOfShifts += 26
		}

		shifted := 'a' + ((s[i] - 'a' + byte(numOfShifts)) % 26)
		result[i] = shifted
	}

	return string(result)
}

func main() {
	// result: "ace"
	// s := "abc"
	// shifts := [][]int{{0,1,0},{1,2,1},{0,2,1}}

	// result: "catz"
	s := "dztz"
	shifts := [][]int{{0,0,0},{1,1,1}}

	// result: ""
	// s := ""
	// shifts := [][]int{}

	result := shiftingLetters(s, shifts)
	fmt.Printf("result = %v\n", result)
}

