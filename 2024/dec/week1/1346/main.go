package main
import (
	"fmt"
)

func checkIfExist(arr []int) bool {
	m := make(map[int]int, len(arr))
	for _, num := range arr {
		m[num]++
	}

	for _, num := range arr {
		if num != 0 && m[2 * num] >= 1 {
			return true
		}

		if num == 0 && m[num] > 1 {
			return true
		}
	}

	return false
}

func main() {
	// result: true
	// arr := []int{10,2,5,3}

	// result: false
	arr := []int{3,1,7,11}

	// result: 
	// arr := []int{}

	result := checkIfExist(arr)
	fmt.Printf("result = %v\n", result)
}

