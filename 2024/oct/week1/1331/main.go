package main
import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	m := make(map[int]int)
	sortedArr := make([]int, n)
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	rank := int(1)

	for i, num := range sortedArr {
		if i > 0 && num > sortedArr[i - 1] {
			rank++
		}

		m[num] = rank
	}

	for i, _ := range arr {
		arr[i] = m[arr[i]]
	}

	return arr
}

func main() {
	// result: [4,1,2,3]
	// arr := []int{40,10,20,30}

	// result:  [1,1,1]
	// arr := []int{100,100,100}

	// result: [5,3,4,2,8,6,7,1,3]
	arr := []int{37,12,28,9,100,56,80,5,12}

	// result: 
	// arr := []int{}

	result := arrayRankTransform(arr)
	fmt.Printf("result = %v\n", result)
}

