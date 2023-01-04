package main
import (
	"fmt"
)

func minimumRounds(tasks []int) int {
	counter := map[int]int{}
	for _, task := range tasks {
		counter[task]++
	}

	result := int(0)
	for _, count := range counter {
		if count == 1 {
			return -1
		}
		if count % 3 == 0 {
			result += count / 3
		} else {
			result += count / 3 + 1
		}
	}

	return result
}

func main() {
	// result: 4
	// tasks := []int{2,2,3,3,2,4,4,4,4,4}

	// result: -1
	tasks := []int{2,3,3}

	// result: 
	// tasks := []int{}

	result := minimumRounds(tasks)
	fmt.Printf("result = %v\n", result)
}

