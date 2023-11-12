package main
import (
	"fmt"
)

func numBusesToDestination(routes [][]int, source int, target int) int {
	n := len(routes)
	m := make(map[int][]int)

    for i := 0; i < n; i++ {
        for j := 0; j < len(routes[i]); j++ {
			busstop := routes[i][j]
			buses := m[busstop]
			buses = append(buses, i)
			m[busstop] = buses
        }
    }

	mBuses := make(map[int]bool)
	mStops := make(map[int]bool)

	queue := make([]int, 0, n)
	result := int(0)

    queue = append(queue, source)
    mStops[source] = true

    for len(queue) > 0 {
		size := len(queue)

        for i := 0; i < size; i++ {
			busstop := queue[0]
			queue = queue[1:]

            if busstop == target {
                return result
            }

			buses := m[busstop]

            for _, bus := range buses {
                if _, ok := mBuses[bus]; ok {
                    continue
                }

                for _, stop := range routes[bus] {
                    if _, ok := mStops[stop]; ok {
                        continue
                    }

                    queue = append(queue, stop)
                    mStops[stop] = true
                }

                mBuses[bus] = true
            }
        }

        result++
    }

    return -1
}

func main() {
	// result: 2
	// routes := [][]int{{1,2,7},{3,6,7}}
	// source := int(1)
	// target := int(6)

	// result: -1
	routes := [][]int{{7,12},{4,5,15},{6},{15,19},{9,12,13}}
	source := int(15)
	target := int(12)

	// result: 
	// routes := [][]int{}
	// source := int(0)
	// target := int(0)

	result := numBusesToDestination(routes, source, target)
	fmt.Printf("result = %v\n", result)
}

