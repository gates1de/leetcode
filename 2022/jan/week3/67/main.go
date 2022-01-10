package main
import (
	"fmt"
)

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		diff := len(a) - len(b)
		for i := 0; i < diff; i++ {
			b = "0" + b
		}
	} else {
		diff := len(b) - len(a)
		for i := 0; i < diff; i++ {
			a = "0" + a
		}
	}

	// fmt.Printf("a = %v, b = %v\n", a, b)
	result := ""
	isAdv := false
	for i := len(a) - 1; i >= 0; i-- {
		aChar := a[i]
		bChar := b[i]
		current := "0"

		if aChar == byte('1') && bChar == byte('1') {
			if isAdv {
				current = "1"
			}
			isAdv = true
		} else if aChar == byte('1') || bChar == byte('1') {
			if isAdv {
				isAdv = true
			} else {
				current = "1"
			}
		} else {
			if isAdv {
				current = "1"
				isAdv = false
			}
		}

		result = current + result
	}

	if isAdv {
		result = "1" + result
	}

	if result == "" {
		return "0"
	}
	return result
}

func main() {
	// result: "100"
	// a := "11"
	// b := "1"

	// result: "10101"
	// a := "1010"
	// b := "1011"

	// result: "0"
	// a := "0"
	// b := "0"

	// result: "1000101000111110001101111001110"
	a := "111011101011110111011111110111"
	b := "1001011011111010001111010111"

	// result: 
	// a := ""
	// b := ""

	result := addBinary(a, b)
	fmt.Printf("result = %v\n", result)
}

