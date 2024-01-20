package main
import (
	"fmt"
)

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	for _, num := range arr {
		counter[num]++
	}

	m := make(map[int]int)
	for _, count := range counter {
		m[count]++
		if m[count] >= 2 {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// arr := []int{1,2,2,1,1,3}

	// result: false
	// arr := []int{1,2}

	// result: true
	arr := []int{-3,0,1,-3,1,1,1,-3,10,0}

	// result: 
	// arr := []int{}

	result := uniqueOccurrences(arr)
	fmt.Printf("result = %v\n", result)
}

