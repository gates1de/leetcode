package main

import (
	"fmt"
	"sort"
	"unicode"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	result := make([]string, 0, len(code))
	business := []string{"electronics", "grocery", "pharmacy", "restaurant"}
	left := int(0)
	right := int(0)
	for _, s := range business {
		for i := range len(code) {
			if isActive[i] && businessLine[i] == s && isValid(code[i]) {
				result = append(result, code[i])
				right++
			}
		}

		sort.Strings(result[left:right])
		left = right
	}

	return result
}

func isValid(code string) bool {
	if len(code) == 0 {
		return false
	}

	for _, r := range code {
		if !unicode.IsNumber(r) && !unicode.IsLetter(r) && r != '_' {
			return false
		}
	}

	return true
}


func main() {
	// result: ["PHARMA5","SAVE20"]
	// code := []string{"SAVE20","","PHARMA5","SAVE@20"}
	// businessLine := []string{"restaurant","grocery","pharmacy","restaurant"}
	// isActive := []bool{true,true,true,true}

	// result: ["ELECTRONICS_50"]
	code := []string{"GROCERY15","ELECTRONICS_50","DISCOUNT10"}
	businessLine := []string{"grocery","electronics","invalid"}
	isActive := []bool{false,true,true}

	// result: []
	// code := []string{}
	// businessLine := []string{}
	// isActive := []bool{}

	result := validateCoupons(code, businessLine, isActive)
	fmt.Printf("result = %v\n", result)
}

