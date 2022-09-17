package main
import (
	"fmt"
	"sort"
)

func findOriginalArray(changed []int) []int {
	n := len(changed)
	if n % 2 != 0 {
		return []int{}
	}

	sort.Ints(changed)

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[changed[i]]++
	}

	result := []int{}
	for i := n - 1; i >= 0; i-- {
		num := changed[i]
		if m[num] == 0 {
			continue
		}
		if num % 2 != 0 || m[num / 2] == 0 {
			return []int{}
		}

		result = append(result, num / 2)

		m[num]--
		m[num / 2]--

		if m[num] == 0 {
			delete(m, num)
		}
		if m[num / 2] == 0 {
			delete(m, num / 2)
		}
	}

	if len(m) > 0 {
		return []int{}
	}

	return result
}

func main() {
	// result: [1,3,4]
	// changed := []int{1,3,4,2,6,8}

	// result: []
	// changed := []int{6,3,0,1}

	// result: []
	// changed := []int{1}

	// result: [1,4]
	changed := []int{1,2,4,8}

	// result: 
	// changed := []int{}

	result := findOriginalArray(changed)
	fmt.Printf("result = %v\n", result)
}

