package main
import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	if isAlwaysFailure(s) {
		return false
	}

	nextCharList := map[string][]string{
		"0": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"1": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"2": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"3": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"4": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"5": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"6": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"7": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"8": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"9": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".", "e", "E"},
		"+": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."},
		"-": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."},
		"e": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "-"},
		"E": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "-"},
		".": {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}
	var splitS []string
	if strings.Contains(s, "e") && strings.Contains(s, "E") {
		return false
	} else if strings.Contains(s, "e") {
		splitS = strings.Split(s, "e")
	} else if strings.Contains(s, "E") {
		splitS = strings.Split(s, "E")
	} else {
		splitS = []string{s}
	}

	if len(splitS) >= 3 {
		return false
	}

	sub := splitS[0]
	if isAlwaysFailure(sub) ||
		strings.Count(sub, ".") >= 2 ||
		strings.Count(sub, "+") >= 2 ||
		strings.Count(sub, "-") >= 2 {
		return false
	}

	nextChars := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "-", "."}
	for _, r := range sub {
		char := string(r)
		if !contains(char, nextChars) {
			return false
		}
		nextChars = nextCharList[char]
	}

	if len(splitS) == 2 {
		sub = splitS[1]
		if len(sub) == 0 || strings.Contains(sub, ".") {
			return false
		}

		if strings.Contains(s, "e") {
			nextChars = nextCharList["e"]
		} else if strings.Contains(s, "E") {
			nextChars = nextCharList["E"]
		}

		for _, r := range sub {
			char := string(r)
			if !contains(char, nextChars) {
				return false
			}
			nextChars = nextCharList[char]
		}
	}

	return true
}

func contains(s string, list []string) bool {
	for _, r := range list {
		if s == string(r) {
			return true
		}
	}
	return false
}

func isAlwaysFailure(s string) bool {
	return len(s) == 0 ||
		contains(s, []string{".", "+", "-", "e", "E", "+.", "-."}) ||
		contains(string(s[len(s) - 1]), []string{"+", "-"})
}

func main() {
	// result: true
	// s := "0"

	// result: false
	// s := "e"

	// result: false
	// s := "."

	// result: true
	// s := ".1"

	// result: true
	// s := "+3.14"

	// result: true
	// s := "-90E3"

	// result: true
	// s := "53.5e93"

	// result: true
	// s := "-123.456e789"

	// result: false
	// s := "abc"

	// result: false
	// s := "1a"

	// result: false
	// s := "1e"

	// result: false
	// s := "e3"

	// result: false
	// s := "99e2.5"

	// result: false
	// s := "-+3"

	// result: false
	// s := "95a54e53"

	// result: false
	// s := "95a54e53"

	// result: false
	// s := ".e1"

	// result: false
	// s := ".1."

	// result: false
	// s := "te1"

	// result: false
	// s := "6+1"

	// result: false
	s := "4e+"

	// result: 
	// s := ""

	result := isNumber(s)
	fmt.Printf("result = %v\n", result)
}

