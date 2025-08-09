package main

import (
	"fmt"
)

func totalFruit(fruits []int) int {
	n := len(fruits)
	basket := make(map[int]int, n)
	left := int(0)
	result := int(0)

	for right := 0; right < n; right++ {
		basket[fruits[right]]++

		for len(basket) > 2 {
			basket[fruits[left]]--
			if basket[fruits[left]] == 0 {
				delete(basket, fruits[left])
			}
			left++
		}

		result = max(result, right - left + 1)
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// fruits := []int{1,2,1}

	// result: 3
	// fruits := []int{0,1,2,2}

	// result: 4
	fruits := []int{1,2,3,2,2}

	// result: 
	// fruits := []int{}

	result := totalFruit(fruits)
	fmt.Printf("result = %v\n", result)
}

