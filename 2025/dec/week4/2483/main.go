package main

import (
	"fmt"
)

func bestClosingTime(customers string) int {
	penalty := int(0)
	for _, customer := range customers {
		if customer == 'Y' {
			penalty++
		}
	}

	minPenalty := penalty
	result := int(-1)

	for i, customer := range customers {
		if customer == 'Y' {
			penalty--
		} else {
			penalty++
		}

		if penalty < minPenalty {
			minPenalty = penalty
			result = i
		}
	}

	if result == -1 {
		return 0
	}

	return result + 1
}

func main() {
	// result: 2
	// customers := "YYNY"

	// result: 0
	// customers := "NNNNN"

	// result: 4
	// customers := "YYYY"

	// result: 1
	customers := "YN"

	// result: 
	// customers := ""

	result := bestClosingTime(customers)
	fmt.Printf("result = %v\n", result)
}

