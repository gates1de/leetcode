package main

import (
	"fmt"
)

func chalkReplacer(chalk []int, k int) int {
	sum := int(0)

	for _, c := range chalk {
		sum += c
	}

	k = k % sum
	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}

	return 0
}

func main() {
	// result: 0
	// chalk := []int{5, 1, 5}
	// k := int(22)

	// result: 1
	chalk := []int{3, 4, 1, 2}
	k := int(25)

	// result:
	// chalk := []int{}
	// k := int(0)

	result := chalkReplacer(chalk, k)
	fmt.Printf("result = %v\n", result)
}
