package main
import (
	"fmt"
	"strings"
)

// Fast
func isValid(s string) bool {
	chars := strings.Split(s, "")
	charsLen := len(chars)
	if charsLen == 0 || charsLen % 2 == 1 {
		return false
	}

	stack := []string{}
	for _, c := range chars {
		// fmt.Printf("chars = %v, c = %v, stack = %v\n", chars, c, stack)
		if strings.Contains("({[", c) {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 && strings.Contains(")}]", c) {
			return false
		}

		target := stack[len(stack) - 1]
		if target == "(" && c == ")" ||
			target == "{" && c == "}" ||
			target == "[" && c == "]" {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}
	if strings.Join(stack, "") == "" {
		return true
	}
	return false
}

// Slow & Use a lot of memory
func mySolution(s string) bool {
	chars := strings.Split(s, "")
	charsLen := len(chars)
	if charsLen == 0 || charsLen % 2 == 1 {
		return false
	}

	for i := charsLen - 2; i >= 0; i-- {
		c := chars[i]
		// fmt.Printf("chars = %v, c = %v\n", chars, c)
		if c == "" || strings.Contains(")}]", c) {
			continue
		}

		j := i + 1
JLOOP:
		for j < len(chars) {
			nextC := chars[j]
			// fmt.Printf("c = %v, nextC = %v\n", c, nextC)
			switch c {
			case "(":
				if nextC == ")" {
					chars[i] = ""
					chars[j] = ""
					break JLOOP
				} else if nextC == "}" || nextC == "]" {
					return false
				}
			case "{":
				if nextC == "}" {
					chars[i] = ""
					chars[j] = ""
					break JLOOP
				} else if nextC == ")" || nextC == "]" {
					return false
				}
			case "[":
				if nextC == "]" {
					chars[i] = ""
					chars[j] = ""
					break JLOOP
				} else if nextC == ")" || nextC == "}" {
					return false
				}
			}
			j++
		}
	}

	if strings.Join(chars, "") == "" {
		return true
	}
	return false
}

func main() {
	// s := "()"
	// s := "()[]{}"
	// s := "(]"
	// s := "([)]"
	// s := "{[]}"
	// s := "(("
	// s := "(]"
	// s := "([}}])"
	s := "){"
	result := isValid(s)
	fmt.Printf("result = %v\n", result)
}

