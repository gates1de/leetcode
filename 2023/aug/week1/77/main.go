package main
import (
	"fmt"
)

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0, k)
    helper(1, n, current, &result)
    return result
}

func helper(start int, limit int, current []int, result *[][]int) {
    if len(current) == cap(current) {
        *result = append(*result, append([]int{}, current...))
        return
    } 

    if len(current) + (limit - start + 1) < cap(current) {
        return
    }

    for i := start; i <= limit; i++ {
        current = append(current, i)
        helper(i + 1, limit, current, result)
        current = current[:len(current) - 1]
    }
}

func main() {
	// result: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
	// n := int(4)
	// k := int(2)

	// result: [[1]]
	n := int(1)
	k := int(1)

	// result: 
	// n := int(0)
	// k := int(0)

	result := combine(n, k)
	fmt.Printf("result = %v\n", result)
}

