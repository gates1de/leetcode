package main
import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	n := len(num1)
	long := num1
	short := num2
	if n < len(num2) {
		n = len(num2)
		long, short = num2, num1
	}

	// fmt.Printf("long = %v, short = %v\n", long, short)
	result := ""
	carryUp := int(0)
	for i := 0; i < n; i++ {
		longNum, _ := strconv.Atoi(string(long[len(long) - 1 - i]))
		shortNum := int(0)
		if len(short) - 1 - i >= 0 {
			shortNum, _ = strconv.Atoi(string(short[len(short) - 1 - i]))
		}
		num := longNum + shortNum + carryUp
		if num >= 10 {
			carryUp = 1
			num -= 10
		} else {
			carryUp = 0
		}
		result = fmt.Sprintf("%v", num) + result
	}

	if carryUp > 0 {
		result = fmt.Sprintf("%v", 1) + result
	}

	return result
}

func main() {
	// result: "134"
	// num1 := "11"
	// num2 := "123"

	// result: "533"
	// num1 := "456"
	// num2 := "77"

	// result: "0"
	// num1 := "0"
	// num2 := "0"

	// result: "10"
	num1 := "1"
	num2 := "9"

	// result: "32527986068802208573"
	// num1 := "8902147328910431"
	// num2 := "32519083921473298142"

	// result: 
	// num1 := ""
	// num2 := ""

	result := addStrings(num1, num2)
	fmt.Printf("result = %v\n", result)
}

