package main

import (
	"fmt"
)

func totalNumbers(digits []int) int {
	count := [10]int{}
	for _, digit := range digits {
		count[digit]++
	}

	result := 0
	for hundreds := 1; hundreds <= 9; hundreds++ {
		if count[hundreds] == 0 {
			continue
		}
		count[hundreds]--
		for tens := 0; tens <= 9; tens++ {
			if count[tens] == 0 {
				continue
			}
			count[tens]--
			for units := 0; units <= 8; units += 2 {
				if count[units] > 0 {
					result++
				}
			}
			count[tens]++
		}
		count[hundreds]++
	}

	return result
}

func main() {
	// result: 12
	// digits := []int{1,2,3,4}

	// result: 2
	// digits := []int{0,2,2}

	// result: 1
	// digits := []int{6,6,6}

	// result: 0
	digits := []int{1, 3, 5}

	// result: 0
	// digits := []int{}

	result := totalNumbers(digits)
	fmt.Printf("result = %v\n", result)
}
