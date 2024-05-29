package main
import (
	"fmt"
	"strconv"
)

func numSteps(s string) int {
	n := len(s)
	ope := int(0)
	carry := int(0)

	for i := n - 1; i > 0; i-- {
		digit, _ := strconv.Atoi(string(s[i]))
		digit += carry

		if digit % 2 == 1 {
			ope += 2
			carry = 1
		} else {
			ope++
		}
	}

	return ope + carry
}

func main() {
	// result: 6
	// s := "1101"

	// result: 1
	// s := "10"

	// result: 0
	s := "1"

	// result: 
	// s := ""

	result := numSteps(s)
	fmt.Printf("result = %v\n", result)
}

