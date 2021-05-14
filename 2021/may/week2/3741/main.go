package main
import (
	"fmt"
	"strings"
)

// Hint: Cartesian Product (直積集合, デカルト積)

func ambiguousCoordinates(s string) []string {
	numString := string(s[1:len(s) - 1])
	result := []string{}
	if len(numString) == 0 {
		return result
	}

	for i, _ := range numString {
		x := string(numString[:i + 1])
		// fmt.Printf("x = %v\n", x)

		if isInt(x) {
			y := numString[i + 1:]
			// fmt.Printf("y = %v\n", y)
			if isInt(y) {
				result = append(result, fmt.Sprintf("(%v, %v)", x, y))
			}

			yWithDot := ""
			for j, _ := range y {
				yWithDot = string(y[:j + 1]) + "." + string(y[j + 1:])
				// fmt.Printf("yWithDot = %v\n", yWithDot)
				if isFloat(yWithDot) {
					result = append(result, fmt.Sprintf("(%v, %v)", x, yWithDot))
				}
			}
		}

		if len(x) < 2 {
			continue
		}
		for j := 0; j < len(x); j++ {
			xWithDot := string(x[:j + 1]) + "." + string(x[j + 1:])
			// fmt.Printf("xWithDot = %v\n", xWithDot)
			if !isFloat(xWithDot) {
				continue
			}

			y := numString[i + 1:]
			if isInt(y) {
				result = append(result, fmt.Sprintf("(%v, %v)", xWithDot, y))
			}

			yWithDot := ""
			// fmt.Printf("y = %v\n", y)
			for k, _ := range y {
				yWithDot = string(y[:k + 1]) + "." + string(y[k + 1:])
				// fmt.Printf("yWithDot = %v\n", yWithDot)
				if isFloat(yWithDot) {
					result = append(result, fmt.Sprintf("(%v, %v)", xWithDot, yWithDot))
				}
			}
		}
	}
	return result
}

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	if string(s[0]) == "0" {
		return false
	}
	return true
}

func isFloat(s string) bool {
	if len(s) <= 2 || isAllZero(s) {
		return false
	}
	dotIndex := strings.Index(s, ".")
	if dotIndex < 0 {
		return false
	}
	intNum := s[:dotIndex]
	deciNum := s[dotIndex + 1:]
	// fmt.Printf("intNum = %v, deciNum = %v\n", intNum, deciNum)
	if !isInt(intNum) {
		return false
	}
	if len(deciNum) == 0 || string(deciNum[len(deciNum) - 1]) == "0" {
		return false
	}
	return true
}

func isAllZero(s string) bool {
	for _, r := range s {
		if r == rune('.') {
			continue
		}
		if r != rune('0') {
			return false
		}
	}
	return true
}

func main() {
	// result: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]
	// s := "(123)"

	// result: ["(0.001, 1)", "(0, 0.011)"]
	// s := "(00011)"

	// result: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"]
	// s := "(0123)"

	// result: ["(10, 0)"]
	s := "(100)"

	// result: 
	// s := ""

	result := ambiguousCoordinates(s)
	fmt.Printf("result = %v\n", result)
}

