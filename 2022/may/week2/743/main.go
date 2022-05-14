package main
import (
	"fmt"
)

func networkDelayTime(times [][]int, n int, k int) int {
	dirs := make([][][2]int, n)
	for _, time := range times {
		node := time[0] - 1
		if dirs[node] == nil {
			dirs[node] = [][2]int{}
		}
		dirs[node] = append(dirs[node], [2]int{time[1] - 1, time[2]})
	}

	m := map[int]int{}
	helper(k - 1, 0, dirs, n, m)
	if len(m) != n {
		return -1
	}

	result := int(0)
	for _, t := range m {
		if result < t {
			result = t
		}
	}

	return result
}

func helper(current int, t int, dirs [][][2]int, n int, m map[int]int) {
	nexts := dirs[current]
	if minT, ok := m[current]; ok && minT <= t {
		return
	}

	m[current] = t

	for _, next := range nexts {
		helper(next[0], t + next[1], dirs, n, m)
	}
}

func main() {
	// result: 2
	// times := [][]int{{2,1,1},{2,3,1},{3,4,1}}
	// n := int(4)
	// k := int(2)

	// result: 1
	// times := [][]int{{1,2,1}}
	// n := int(2)
	// k := int(1)

	// result: -1
	// times := [][]int{{1,2,1}}
	// n := int(2)
	// k := int(2)

	// result: 2
	// times := [][]int{{2,1,1},{2,3,1},{1,4,2},{3,4,1}}
	// n := int(4)
	// k := int(2)

	// result: 3
	times := [][]int{{2,1,1},{2,3,3},{1,4,2},{3,4,1}}
	n := int(4)
	k := int(2)

	// result: 
	// times := [][]int{}
	// n := int(0)
	// k := int(0)

	result := networkDelayTime(times, n, k)
	fmt.Printf("result = %v\n", result)
}

