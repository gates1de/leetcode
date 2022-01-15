package main
import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	result := int(0)
	isAppearedNum := false
	isAppearedPlus := false
	isAppearedMinus := false
	for _, char := range s {
		b := byte(char)
		if byte('0') <= b && b <= byte('9') {
			result *= 10
			result += byteToInt(b)
			isAppearedNum = true

			if math.MaxInt32 <= result {
				break
			}
		} else if b == byte('-') {
			if isAppearedNum || isAppearedPlus || isAppearedMinus {
				break
			}
			isAppearedMinus = true
		} else if b == byte('.') {
			break
		} else if b == byte('+') {
			if isAppearedNum || isAppearedPlus || isAppearedMinus {
				break
			}
			isAppearedPlus = true
		} else if b == byte(' ') {
			if isAppearedNum || isAppearedPlus || isAppearedMinus {
				break
			}
		} else {
			if isAppearedNum {
				break
			}
			return 0
		}
	}

	if isAppearedMinus {
		result *= -1
		if result <= math.MinInt32 {
			return math.MinInt32
		}
		return result
	}

	if math.MaxInt32 <= result {
		return math.MaxInt32
	}
	return result
}

func byteToInt(b byte) int {
	bytes := map[byte]int{
		byte('0'): 0,
		byte('1'): 1,
		byte('2'): 2,
		byte('3'): 3,
		byte('4'): 4,
		byte('5'): 5,
		byte('6'): 6,
		byte('7'): 7,
		byte('8'): 8,
		byte('9'): 9,
	}
	return bytes[b]
}

func main() {
	// result: 42
	// s := "42"

	// result: -42
	// s := "   -42"

	// result: 4193
	// s := "4193 with words"

	// result: 2147483647
	// s := "2147483648"

	// result: -2147483648
	// s := "-2147483649"

	// result: 1
	// s := "1-2"

	// result: 0
	// s := "--2"

	// result: 42
	// s := "42.5"

	// result: 0
	// s := ".42"

	// result: 42
	// s := "42+5"

	// result: 0
	// s := "words and 987"

	// result: -42
	// s := "   -42   8"

	// result: 0
	// s := "+-12"

	// result: 0
	// s := "-+12"

	// result: 4
	// s := "-4+12"

	// result: 2147483647
	// s := "9223372036854775808"

	// result: 0
	s := "  +  413"

	// result: 
	// s := ""

	result := myAtoi(s)
	fmt.Printf("result = %v\n", result)
}

