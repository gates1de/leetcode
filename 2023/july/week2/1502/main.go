package main
import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	n := len(arr)
	if n <= 1 {
		return true
	}

	sort.Ints(arr)
	startDiff := arr[1] - arr[0] 
	for i := 1; i < n - 1; i++ {
		diff := arr[i + 1] - arr[i]
		if startDiff != diff {
			return false
		}
	}
	return true
}

func main() {
	// result: true
	// arr := []int{3,5,1}

	// result: false
	// arr := []int{1,2,4}

	// result: true
	arr := []int{1,1}

	// result: 
	// arr := []int{}

	result := canMakeArithmeticProgression(arr)
	fmt.Printf("result = %v\n", result)
}

