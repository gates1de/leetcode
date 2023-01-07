package main
import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(a, b int) bool {
        if points[a][0] == points[b][0] {
            return points[a][1] < points[b][1]
        }
        return points[a][0] < points[b][0]
    })

    result := int(1)
    currentEnd := points[0][1]
    for _, point := range points[1:] {
        start := point[0]
        end := point[1]

        if currentEnd == -1 {
            result++
            currentEnd = end
            continue
        }

        if currentEnd <= start {
            if currentEnd == start {
                currentEnd = start
            } else {
                result++
                currentEnd = end
            }
        } else {
            currentEnd = min(currentEnd, end)
        }
    }

    return result
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func main() {
	// result: 2
	// points := [][]int{{10,16},{2,8},{1,6},{7,12}}

	// result: 4
	// points := [][]int{{1,2},{3,4},{5,6},{7,8}}

	// result: 2
	points := [][]int{{1,2},{2,3},{3,4},{4,5}}

	// result: 
	// points := [][]int{}

	result := findMinArrowShots(points)
	fmt.Printf("result = %v\n", result)
}

