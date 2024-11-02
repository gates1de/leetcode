package main
import (
	"fmt"
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(a, b int) bool { return factory[a][0] < factory[b][0] })

	positions := make([]int, 0)
	for _, f := range factory {
		for i := 0; i < f[1]; i++ {
			positions = append(positions, f[0])
		}
	}

	robotCount := len(robot)
	factoryCount := len(positions)

	next := make([]int64, factoryCount + 1)
	current := make([]int64, factoryCount + 1)

	for i := robotCount - 1; i >= 0; i-- {
		if i != robotCount - 1 {
			next[factoryCount] = int64(1e12)
		}

		current[factoryCount] = int64(1e12)

		for j := factoryCount - 1; j >= 0; j-- {
			assign := int64(abs(robot[i] - positions[j])) + next[j + 1]
			skip := current[j + 1]
			current[j] = min(assign, skip)
		}

		copy(next, current)
	}

	return current[0]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int64) int64 {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// robot := []int{0,4,6}
	// factory := [][]int{{2,2},{6,2}}

	// result: 2
	robot := []int{1,-1}
	factory := [][]int{{-2,1},{2,1}}

	// result: 
	// robot := []int{}
	// factory := [][]int{}

	result := minimumTotalDistance(robot, factory)
	fmt.Printf("result = %v\n", result)
}

