package main
import (
	"fmt"
)

func intToRoman(num int) string {
	result := ""
	for num > 0 {
		if num >= 1000 {
			result += "M"
			num -= 1000
		} else if 900 <= num && num <= 999 {
			result += "CM"
			num -= 900
		} else if 500 <= num && num <= 899 {
			result += "D"
			num -= 500
		} else if 400 <= num && num <= 499 {
			result += "CD"
			num -= 400
		} else if 100 <= num && num <= 399 {
			result += "C"
			num -= 100
		} else if 90 <= num && num <= 99 {
			result += "XC"
			num -= 90
		} else if 50 <= num && num <= 89 {
			result += "L"
			num -= 50
		} else if 40 <= num && num <= 49 {
			result += "XL"
			num -= 40
		} else if 10 <= num && num <= 39 {
			result += "X"
			num -= 10
		} else if num == 9 {
			result += "IX"
			num -= 9
		} else if 5 <= num && num <= 8 {
			result += "V"
			num -= 5
		} else if num == 4 {
			result += "IV"
			num -= 4
		} else {
			result += "I"
			num -= 1
		}
	}

	return result
}

func main() {
	// result: "III"
	// num := int(3)

	// result: "LVIII"
	// num := int(58)

	// result: "MCMXCIV"
	num := int(1994)

	// result: 
	// num := int(0)

	result := intToRoman(num)
	fmt.Printf("result = %v\n", result)
}

