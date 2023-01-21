package main
import (
	"fmt"
)

func restoreIpAddresses(s string) []string {
	n := len(s)
	result := make([]string, 0, 10000)
	if n <= 3 {
		return result
	}
	for a := 1; a <= min(3, n - 1); a++ {
		for b := a + 1; b <= min(a + 3, n - 1); b++ {
			for c := b + 1; c <= min(b + 3, n - 1); c++ {
				classA := s[:a]
				classB := s[a:b]
				classC := s[b:c]
				classD := s[c:]
				if len(classC) >= 4 || len(classD) >= 4 {
					continue
				}
				isValidOctetA := isValidOctet(classA)
				isValidOctetB := isValidOctet(classB)
				isValidOctetC := isValidOctet(classC)
				isValidOctetD := isValidOctet(classD)
				if isValidOctetA && isValidOctetB && isValidOctetC && isValidOctetD {
					result = append(result, classA + "." + classB + "." + classC + "." + classD)
				}
			}
		}
	}
	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func isValidOctet(octet string) bool {
	switch len(octet) {
	case 1:
		return true
	case 2:
		return octet[0] != '0'
	case 3:
		if octet[0] == '0' {
			return false
		}

		if octet[0] >= '3' {
			return false
		}

		if octet[0] == '2' {
			if octet[1] >= '6' {
				return false
			}
			if octet[1] == '5' {
				return octet[2] <= '5'
			}
			return true
		}
		return true
	default:
		return false
	}
}

func main() {
	// result: ["255.255.11.135","255.255.111.35"]
	// s := "25525511135"

	// result: ["0.0.0.0"]
	// s := "0000"

	// result: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
	s := "101023"

	// result: 
	// s := ""

	result := restoreIpAddresses(s)
	fmt.Printf("result = %v\n", result)
}

