package main
import (
	"fmt"
)

func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    result := make([][]int, 0, n)
    newStart := newInterval[0]
    newEnd := newInterval[1]
    isInserted := false

    for _, interval := range intervals {
        start := interval[0]
        end := interval[1]

        if end < newStart {
            result = append(result, interval)
            continue
        }

        if newEnd < start {
            if !isInserted {
                result = append(result, []int{newStart, newEnd})
                isInserted = true
            }
            result = append(result, interval)
            continue
        }

        if !isInserted && newEnd == start {
            result = append(result, []int{newStart, end})
            isInserted = true
            continue
        }

        if start <= newStart && newStart <= end {
            newStart = min(newStart, start)
        }
        if start <= newEnd && newEnd <= end {
            newEnd = max(newEnd, end)
        }
    }

    if !isInserted {
        result = append(result, []int{newStart, newEnd})
    }
    return result
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func main() {
	// result: [[1,5],[6,9]]
	// intervals := [][]int{{1,3},{6,9}}
	// newInterval := []int{2,5}

	// result: [[1,2],[3,10],[12,16]]
	intervals := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	newInterval := []int{4,8}

	// result: 
	// intervals := [][]int{}
	// newInterval := []int{}

	result := insert(intervals, newInterval)
	fmt.Printf("result = %v\n", result)
}

