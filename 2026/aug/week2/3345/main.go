package main

import (
	"fmt"
	"strconv"
)

const maximumNumber = int(100)

func smallestNumber(n int, t int) int {
	result := int(0)
	for offset := range maximumNumber - n + 1 {
		candidate := n + offset
		product := int(1)
		for _, digit := range strconv.Itoa(candidate) {
			product *= int(digit - '0')
		}

		if product % t == 0 {
			result = candidate
			break
		}
	}

	return result
}

func main() {
	// result: 10
	// n := int(10)
	// t := int(2)

	// result: 16
	n := int(15)
	t := int(3)

	// result:
	// n := int(0)
	// t := int(0)

	result := smallestNumber(n, t)
	fmt.Printf("result = %v\n", result)
}
