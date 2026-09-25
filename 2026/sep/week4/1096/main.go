package main

import (
	"fmt"
	"sort"
)

func braceExpansionII(expression string) []string {
	position := 0
	var parseExpression func() map[string]struct{}
	var parseTerm func() map[string]struct{}
	var parseFactor func() map[string]struct{}

	parseFactor = func() map[string]struct{} {
		if expression[position] == '{' {
			position++
			result := parseExpression()
			position++
			return result
		}

		word := string(expression[position])
		position++
		return map[string]struct{}{word: {}}
	}

	parseTerm = func() map[string]struct{} {
		result := map[string]struct{}{"": {}}
		for position < len(expression) && expression[position] != ',' && expression[position] != '}' {
			factor := parseFactor()
			next := make(map[string]struct{}, len(result)*len(factor))
			for left := range result {
				for right := range factor {
					next[left+right] = struct{}{}
				}
			}
			result = next
		}
		return result
	}

	parseExpression = func() map[string]struct{} {
		result := make(map[string]struct{})
		for {
			for word := range parseTerm() {
				result[word] = struct{}{}
			}
			if position >= len(expression) || expression[position] != ',' {
				break
			}
			position++
		}
		return result
	}

	if len(expression) == 0 {
		return []string{}
	}

	words := parseExpression()
	result := make([]string, 0, len(words))
	for word := range words {
		result = append(result, word)
	}

	sort.Strings(result)
	return result
}

func main() {
	// result: ["ac","ad","ae","bc","bd","be"]
	// expression := "{a,b}{c,{d,e}}"

	// result: ["a","ab","ac","z"]
	expression := "{{a,z},a{b,c},{ab,z}}"

	// result: []
	// expression := ""

	result := braceExpansionII(expression)
	fmt.Printf("result = %v\n", result)
}
