package main

import (
	"fmt"
)

func numberOfBeams(bank []string) int {
	result := int(0)
	prev := int(0)

	for _, s := range bank {
		count := int(0)

		for _, char := range s {
			if char == '1' {
				count++
			}
		}

		if count > 0 {
			result += prev * count
			prev = count
		}
	}

	return result
}

func main() {
	// result: 8
	// bank := []string{"011001","000000","010100","001000"}

	// result: 0
	bank := []string{"000","111","000"}

	// result: 
	// bank := []string{}

	result := numberOfBeams(bank)
	fmt.Printf("result = %v\n", result)
}

