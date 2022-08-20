package main
import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := int(0)
	for strings.Index(s, "IV") >= 0 {
		result += 4
		s = strings.Replace(s, "IV", "", 1)
	}
	for strings.Index(s, "IX") >= 0 {
		result += 9
		s = strings.Replace(s, "IX", "", 1)
	}
	for strings.Index(s, "XL") >= 0 {
		result += 40
		s = strings.Replace(s, "XL", "", 1)
	}
	for strings.Index(s, "XC") >= 0 {
		result += 90
		s = strings.Replace(s, "XC", "", 1)
	}
	for strings.Index(s, "CD") >= 0 {
		result += 400
		s = strings.Replace(s, "CD", "", 1)
	}
	for strings.Index(s, "CM") >= 0 {
		result += 900
		s = strings.Replace(s, "CM", "", 1)
	}

	for _, char := range s {
		result += m[char]
	}
	return result
}

func main() {
	// result: 3
	// s := "III"

	// result: 58
	// s := "LVIII"

	// result: 1994
	s := "MCMXCIV"

	// result: 
	// s := ""

	result := romanToInt(s)
	fmt.Printf("result = %v\n", result)
}

