package main
import (
	"fmt"
)

func canArrange(arr []int, k int) bool {
	m := make(map[int]int)

	for _, num := range arr {
		remain := ((num % k) + k) % k
		m[remain]++
	}

	for _, num := range arr {
		remain := ((num % k) + k) % k
		if remain == 0 {
			if m[remain] % 2 == 1 {
				return false
			}
		} else if m[remain] != m[k - remain] {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// arr := []int{1,2,3,4,5,10,6,7,8,9}
	// k := int(5)

	// result: true
	// arr := []int{1,2,3,4,5,6}
	// k := int(7)

	// result: false
	arr := []int{1,2,3,4,5,6}
	k := int(10)

	// result: 
	// arr := []int{}
	// k := int(0)

	result := canArrange(arr, k)
	fmt.Printf("result = %v\n", result)
}

