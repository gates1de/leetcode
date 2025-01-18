package main
import (
	"fmt"
)

func canBeValid(s string, locked string) bool {
	n := len(s)

	if n % 2 == 1 {
		return false
	}

	openBrackets := int(0)
	unlocked := int(0)


	for i, char := range locked {
		if char == '0' {
			unlocked++
		} else if s[i] == '(' {
			openBrackets++
		} else if s[i] == ')' {
			if openBrackets > 0 {
				openBrackets--
			} else if unlocked > 0 {
				unlocked--
			} else {
				return false
			}
		}
	}

	balance := int(0)

	for i := n - 1; i >= 0; i-- {
		if locked[i] == '0' {
			balance--
			unlocked--
		} else if s[i] == '(' {
			balance++
			openBrackets--
		} else if s[i] == ')' {
			balance--
		}

		if balance > 0 {
			return false
		}

		if unlocked == 0 && openBrackets == 0 {
			break
		}
	}

	if openBrackets > 0 {
		return false
	}

	return true
}

func main() {
	// result: true
	// s := "))()))"
	// locked := "010100"

	// result: true
	// s := "()()"
	// locked := "0000"

	// result: false
	s := ")"
	locked := "0"

	// result: 
	// s := ""
	// locked := ""

	result := canBeValid(s, locked)
	fmt.Printf("result = %v\n", result)
}

