package main
import (
	"fmt"
)

func findJudge(n int, trust [][]int) int {
    result := int(-1)
    if n - len(trust) >= 2 {
        return result
    }

    m := make(map[int]int, n + 1)
    for i := 1; i <= n; i++ {
        m[i] = 0
    }

    for _, t := range trust {
        if _, ok := m[t[1]]; ok {
            m[t[1]]++
        }
        delete(m, t[0])
    }

    if len(m) != 1 {
        return result
    }

    for key, val := range m {
        if val != n - 1 {
            return -1
        }
        result = key
    }
    return result
}

func main() {
	// result: 2
	// n := int(2)
	// trust := [][]int{{1,2}}

	// result: 3
	// n := int(3)
	// trust := [][]int{{1,3},{2,3}}

	// result: -1
	n := int(3)
	trust := [][]int{{1,3},{2,3},{3,1}}

	// result: 
	// n := int(0)
	// trust := [][]int{}

	result := findJudge(n, trust)
	fmt.Printf("result = %v\n", result)
}

