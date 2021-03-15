package main
import (
	"fmt"
)

func romanToInt(s string) int {
	romans := map[string]int{
		"I": 1,
		"IV": 4,
		"V": 5,
		"IX": 9,
		"X": 10,
		"XL": 40,
		"L": 50,
		"XC": 90,
		"C": 100,
		"CD": 400,
		"D": 500,
		"CM": 900,
		"M": 1000,
	}
	result := int(0)
	preRomanC := "-"
	for _, romanR := range s {
		romanC := string(romanR)
		twoRomanC := preRomanC + romanC
		if romans[twoRomanC] != 0 {
			result += romans[twoRomanC] - romans[preRomanC]
		} else {
			result += romans[romanC]
		}
		preRomanC = romanC
	}
	return result
}

func main() {
	// s := "III"
	// s := "IV"
	// s := "IX"
	// s := "LVIII"
	s := "MCMXCIV"

	result := romanToInt(s)
	fmt.Printf("result = %v\n", result)
}

