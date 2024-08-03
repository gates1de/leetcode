package main
import (
	"fmt"
)

func canBeEqual(target []int, arr []int) bool {
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}

	for _, num := range target {
		if _, exists := count[num]; !exists {
			return false
		}

		count[num]--
		if count[num] == 0 {
			delete(count, num)
		}
	}

	return len(count) == 0
}

func main() {
	// result: true
	// target := []int{1,2,3,4}
	// arr := []int{2,4,1,3}

	// result: true
	// target := []int{7}
	// arr := []int{7}

	// result: false
	target := []int{3,7,9}
	arr := []int{3,7,11}

	// result: 
	// target := []int{}
	// arr := []int{}

	result := canBeEqual(target, arr)
	fmt.Printf("result = %v\n", result)
}

