package main
import (
	"fmt"
)

func countTriplets(arr []int) int {
	n := len(arr)
	result := int(0)
	prefix := int(0)

	countMap := make(map[int]int)
	totalMap := make(map[int]int)
	countMap[0] = 1

	for i := 0; i < n; i++ {
		prefix ^= arr[i]

		result += countMap[prefix] * i - totalMap[prefix]
		totalMap[prefix] += i + 1
		countMap[prefix]++
	}

	return result
}

func main() {
	// result: 4
	// arr := []int{2,3,1,6,7}

	// result: 10
	arr := []int{1,1,1,1,1}

	// result: 
	// arr := []int{}

	result := countTriplets(arr)
	fmt.Printf("result = %v\n", result)
}

