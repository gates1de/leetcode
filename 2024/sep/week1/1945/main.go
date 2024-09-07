package main

import (
	"fmt"
	"strconv"
)

func getLucky(s string, k int) int {
	strNums := ""
	for _, c := range s {
		strNums += strconv.Itoa(int(c) - 'a' + 1)
	}

	for i := 0; i < k; i++ {
		strNums = strconv.Itoa(getSum(strNums))
	}

	result, _ := strconv.Atoi(strNums)
	return result
}

func getSum(strNums string) int {
	result := int(0)
	for _, c := range strNums {
		result += int(c) - '0'
	}

	return result
}

func main() {
	// result: 36
	// s := "iiii"
	// k := int(1)

	// result: 6
	// s := "leetcode"
	// k := int(2)

	// result: 8
	s := "zbax"
	k := int(2)

	// result:
	// s := ""
	// k := int(0)

	result := getLucky(s, k)
	fmt.Printf("result = %v\n", result)
}
