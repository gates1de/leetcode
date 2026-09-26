package main

import (
	"fmt"
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	values := make(map[string]string, len(knowledge))
	for _, pair := range knowledge {
		values[pair[0]] = pair[1]
	}

	var result strings.Builder
	result.Grow(len(s))
	for i := 0; i < len(s); {
		if s[i] != '(' {
			result.WriteByte(s[i])
			i++
			continue
		}

		start := i + 1
		end := start
		for s[end] != ')' {
			end++
		}
		if value, ok := values[s[start:end]]; ok {
			result.WriteString(value)
		} else {
			result.WriteByte('?')
		}
		i = end + 1
	}

	return result.String()
}

func main() {
	// result: "bobistwoyearsold"
	// s := "(name)is(age)yearsold"
	// knowledge := [][]string{{"name","bob"},{"age","two"}}

	// result: "hi?"
	// s := "hi(name)"
	// knowledge := [][]string{{"a","b"}}

	// result: "yesyesyesaaa"
	s := "(a)(a)(a)aaa"
	knowledge := [][]string{{"a", "yes"}}

	// result: ""
	// s := ""
	// knowledge := [][]string{}

	result := evaluate(s, knowledge)
	fmt.Printf("result = %v\n", result)
}
