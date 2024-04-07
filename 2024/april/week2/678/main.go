package main
import (
	"fmt"
)

func checkValidString(s string) bool {
	openCount := int(0)
	closeCount := int(0)
	n := len(s) - 1

	for i := 0; i <= n; i++ {
		if s[i] == '(' || s[i] == '*' {
			openCount++
		} else {
			openCount--
		}

		if s[n - i] == ')' || s[n - i] == '*' {
			closeCount++
		} else {
			closeCount--
		}

		if openCount < 0 || closeCount < 0 {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// s := "()"

	// result: true
	// s := "(*)"

	// result: true
	s := "(*))"

	// result: 
	// s := ""

	result := checkValidString(s)
	fmt.Printf("result = %v\n", result)
}

