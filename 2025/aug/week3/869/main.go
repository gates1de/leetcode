package main

import (
	"fmt"
	"reflect"
	"sort"
)

func reorderedPowerOf2(n int) bool {
	if n == 0 {
		return false
	}

	digits := numToDigits(n)
	sort.Ints(digits)

	for i := 1; i <= int(10e+9); i *= 2 {
		twoDigits := numToDigits(i)
		sort.Ints(twoDigits)
		if reflect.DeepEqual(digits, twoDigits) {
			return true
		}
	}

	return false
}

func numToDigits(num int) []int {
	result := []int{}
	for remain := num; remain > 0; remain /= 10 {
		result = append(result, remain % 10)
	}
	return result
}

func main() {
	// result: true
	// n := int(1)

	// result: false
	// n := int(10)

	// result: true
	// n := int(821)

	// result: false
	// n := int(8388602)

	// result: true
	n := int(8888630)

	// result: 
	// n := int(0)

	result := reorderedPowerOf2(n)
	fmt.Printf("result = %v\n", result)
}

