package main
import (
	"fmt"
)

func findSpecialInteger(arr []int) int {
	counter := make(map[int]int)
	for _, num := range arr {
		counter[num]++
		if float32(counter[num]) / float32(len(arr)) > 0.25 {
			return num
		}
	}

	return -1
}

func main() {
	// result: 6
	// arr := []int{1,2,2,6,6,6,6,7,10}

	// result: 1
	arr := []int{1,1}

	// result: 
	// arr := []int{}

	result := findSpecialInteger(arr)
	fmt.Printf("result = %v\n", result)
}

