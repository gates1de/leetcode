package main
import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	distances := make([]int, n)
	for i, _ := range distances {
		distances[i] = math.MaxInt32
	}
	distances[src] = 0
	for i := 0; i <= k; i++ {
		temp := make([]int, n)
		copy(temp, distances)
		for _, flight := range flights {
			if distances[flight[0]] != math.MaxInt32 {
				temp[flight[1]] = min(temp[flight[1]], distances[flight[0]] + flight[2])
			}
		}
		distances = temp
	}
	if distances[dst] == math.MaxInt32 {
		return -1
	}
	return distances[dst]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 700
	// n := int(4)
	// flights := [][]int{{0,1,100},{1,2,100},{2,0,100},{1,3,600},{2,3,200}}
	// src := int(0)
	// dst := int(3)
	// k := int(1)

	// result: 200
	// n := int(3)
	// flights := [][]int{{0,1,100},{1,2,100},{0,2,500}}
	// src := int(0)
	// dst := int(2)
	// k := int(1)

	// result: 500
	n := int(3)
	flights := [][]int{{0,1,100},{1,2,100},{0,2,500}}
	src := int(0)
	dst := int(2)
	k := int(0)

	// result: 
	// n := int(0)
	// flights := [][]int{}
	// src := int(0)
	// dst := int(0)
	// k := int(0)

	result := findCheapestPrice(n, flights, src, dst, k)
	fmt.Printf("result = %v\n", result)
}

