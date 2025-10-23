package main

import (
	"fmt"
)

func hasSameDigits(s string) bool {
	n := len(s)
	arr := []byte(s)
	for i := 1; i <= n - 2; i++ {
		for j := 0; j <= n - 1 - i; j++ {
			arr[j] = byte(((int(arr[j] - '0') + int(arr[j + 1] - '0')) % 10) + '0')
		}
	}

	return arr[0] == arr[1]
}

func main() {
	// result: true
	// s := "3902"

	// result: false
	s := "34789"

	// result: 
	// s := ""

	result := hasSameDigits(s)
	fmt.Printf("result = %v\n", result)
}

