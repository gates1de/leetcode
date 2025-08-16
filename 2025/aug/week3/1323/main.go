package main

import (
	"fmt"
)

func maximum69Number(num int) int {
	temp := num
	for i := 1000; i >= 1; i /= 10 {
		digit := temp / i
		if digit == 6 {
			return num + 3 * i
		}
		temp %= i
	}
	return num
}

func main() {
	// result: 9969
	// num := int(9669)

	// result: 9999
	// num := int(9996)

	// result: 9999
	// num := int(9999)

	// result: 966
	// num := int(666)

	// result: 9
	num := int(6)

	// result: 
	// num := int(0)

	result := maximum69Number(num)
	fmt.Printf("result = %v\n", result)
}

