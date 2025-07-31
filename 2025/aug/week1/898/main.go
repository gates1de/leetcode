package main

import (
	"fmt"
)

func subarrayBitwiseORs(arr []int) int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	m2[0] = true

	for _, num1 := range arr {
		tmp := make(map[int]bool)

		for num2 := range m2 {
			tmp[num1 | num2] = true
		}

		tmp[num1] = true
		m2 = tmp
		for num := range m2 {
			m1[num] = true
		}
	}

	return len(m1)
}

func main() {
	// result: 1
	// arr := []int{0}

	// result: 3
	// arr := []int{1,1,2}

	// result: 6
	arr := []int{1,2,4}

	// result: 
	// arr := []int{}

	result := subarrayBitwiseORs(arr)
	fmt.Printf("result = %v\n", result)
}

