package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 || denominator == 0 {
		return "0"
	}

	sign := int(1)
	if numerator < 0 {
		sign *= -1
		numerator *= -1
	}

	if denominator < 0 {
		sign *= -1
		denominator *= -1
	}

	prefix := ""
	if sign < 0 {
		prefix = "-"
	}

	integerPart := numerator/denominator
	decimalPart := numerator%denominator
	integerPartInString := strconv.Itoa(integerPart)

	if decimalPart == 0 {
		return prefix + integerPartInString
	}

	currentDecimal := decimalPart
	order := int(1)
	decimalPartInBytes := make([]byte, 0)
	m := make(map[int]int)

	for currentDecimal != 0 {
		if m[currentDecimal] != 0 {
			repeatHeadPosition := m[currentDecimal] - 1
			repeatPart := make([]byte, order - m[currentDecimal])
			copy(repeatPart, decimalPartInBytes[repeatHeadPosition:])
			decimalPartInBytes = decimalPartInBytes[:repeatHeadPosition]
			decimalPartInBytes = append(decimalPartInBytes, '(')
			decimalPartInBytes = append(decimalPartInBytes, repeatPart...)
			decimalPartInBytes = append(decimalPartInBytes, ')')
			break
		}

		m[currentDecimal] = order
		decimalPartInBytes = append(decimalPartInBytes, byte(currentDecimal * 10 / denominator) + '0')
		currentDecimal = currentDecimal * 10 % denominator
		order++
	}

	return prefix + integerPartInString + "." + string(decimalPartInBytes)
}

func main() {
	// result: "0.5"
	// numerator := int(1)
	// denominator := int(2)

	// result: "2"
	// numerator := int(2)
	// denominator := int(1)

	// result: "0.(012)"
	numerator := int(4)
	denominator := int(333)

	// result: ""
	// numerator := int(0)
	// denominator := int(0)

	result := fractionToDecimal(numerator, denominator)
	fmt.Printf("result = %v\n", result)
}

