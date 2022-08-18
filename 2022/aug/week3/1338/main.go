package main
import (
	"fmt"
	"sort"
)

func minSetSize(arr []int) int {
	m := map[int]int{}
	size := len(arr)
	for _, num := range arr {
		m[num]++
	}

	i := int(0)
	counts := make([]int, len(m))
	for _, count := range m {
		counts[i] = count
		i++
	}
	sort.Slice(counts, func(a, b int) bool {
		return counts[a] > counts[b]
	})

	sum := int(0)
	result := int(0)
	for (size / 2) > sum {
		sum += counts[result]
		result++
	}

	return result
}

func main() {
	// result: 2
	// arr := []int{3,3,3,3,5,5,5,2,2,7}

	// result: 1
	// arr := []int{7,7,7,7,7,7}

	// result: 5
	// arr := []int{1,2,3,4,5,6,7,8,9,10}

	// result: 4
	arr := []int{1,1,2,2,3,4,5,6,7,8,9,10}

	// result: 
	// arr := []int{}

	result := minSetSize(arr)
	fmt.Printf("result = %v\n", result)
}

