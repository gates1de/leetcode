package main
import (
	"fmt"
	"reflect"
	"sort"
)

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	targetTime := times[targetFriend]
	sort.Slice(times, func (a, b int) bool { return times[a][0] < times[b][0] })
	chairTime := make([]int, n)

	for _, time := range times {
		for i := 0; i < n; i++ {
			if chairTime[i] <= time[0] {
				chairTime[i] = time[1]
				if reflect.DeepEqual(time, targetTime) {
					return i
				}
				break
			}
		}
	}

	return 0
}

func main() {
	// result: 1
	// times := [][]int{{1,4},{2,3},{4,6}}
	// targetFriend := int(1)

	// result: 2
	times := [][]int{{3,10},{1,5},{2,6}}
	targetFriend := int(0)

	// result: 
	// times := [][]int{}
	// targetFriend := int(0)

	result := smallestChair(times, targetFriend)
	fmt.Printf("result = %v\n", result)
}

