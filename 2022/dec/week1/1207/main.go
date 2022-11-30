package main
import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	sort.Ints(arr)
	m := make(map[int]bool, 1000)
	pre := int(-1001)
	count := int(0)
	for _, num := range arr {
		if pre == num {
			count++
		} else {
			if m[count] {
				return false
			}
			m[count] = true
			count = 1
		}
		pre = num
	}

	if m[count] {
		return false
	}
	m[count] = true

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

