package main
import (
	"fmt"
	"sort"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	starts := make([]int, 0, len(flowers))
	ends := make([]int, 0, len(flowers))

	for _, flower := range flowers {
		starts = append(starts, flower[0])
		ends = append(ends, flower[1] + 1)
	}

	sort.Ints(starts)
	sort.Ints(ends)
	result := make([]int, len(people))

	for index, person := range people {
		i := binarySearch(starts, person)
		j := binarySearch(ends, person)
		result[index] = i - j
	}

	return result
}

func binarySearch(nums []int, target int) int {
	left := int(0)
	right := len(nums)

	for left < right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	// result: [1,2,2,2]
	// flowers := [][]int{{1,6},{3,7},{9,12},{4,13}}
	// people := []int{2,3,7,11}

	// result: [2,2,1]
	flowers := [][]int{{1,10},{3,3}}
	people := []int{3,3,2}

	// result: 
	// flowers := [][]int{}
	// people := []int{}

	result := fullBloomFlowers(flowers, people)
	fmt.Printf("result = %v\n", result)
}

