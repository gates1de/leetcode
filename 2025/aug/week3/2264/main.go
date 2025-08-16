package main

import (
	"fmt"
)

func largestGoodInteger(num string) string {
	result := ""
	if len(num) <= 2 {
		return ""
	}

	nums := make(map[rune]int, 10)
	pre := rune(num[0])
	nums[pre]++

	for _, char := range num[1:] {
		if pre != char {
			nums = make(map[rune]int, 10)
		}

		nums[char]++

		if nums[char] >= 3 && result < string(char) {
			result = string(char)
		}

		pre = char
	}

	return result + result + result
}

func main() {
	// result: "777"
	// num := "6777133339"

	// result: "000"
	// num := "2300019"

	// result: ""
	// num := "42352338"

	// result: "222"
	num := "222"

	// result: ""
	// num := ""

	result := largestGoodInteger(num)
	fmt.Printf("result = %v\n", result)
}

