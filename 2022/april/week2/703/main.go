package main
import (
	"fmt"
	"sort"
)

type KthLargest struct {
	K int
	SortedNums []int
}

func Constructor(k int, nums []int) KthLargest {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	return KthLargest{k, sortedNums}
}

func (this *KthLargest) Add(val int) int {
	n := len(this.SortedNums)
	index := sort.SearchInts(this.SortedNums, val)
	newSortedNums := make([]int, n + 1)
	copy(newSortedNums, this.SortedNums[:index])
	newSortedNums[index] = val
	copy(newSortedNums[index + 1:], this.SortedNums[index:])
	this.SortedNums = newSortedNums
	n = len(this.SortedNums)
	return this.SortedNums[n - this.K]
}

func main() {
	// result: [null, 4, 5, 5, 8, 8]
	// nums := []int{4, 5, 8, 2}
	// k := int(3)
	// addedNums := []int{3, 5, 10, 9, 4}

	// result: [null, 1, 1, 1, 1, 1]
	// nums := []int{1, 2, 3, 4, 5}
	// k := int(5)
	// addedNums := []int{-1, -2, -3, -4, -5}

	// result: [null, 1]
	nums := []int{}
	k := int(1)
	addedNums := []int{}

	// result: 
	// nums := []int{}
	// k := int(0)
	// addedNums := []int{}

	obj := Constructor(k, nums)
	for _, num := range addedNums {
		fmt.Printf("obj.Add(%v) = %v\n", num, obj.Add(num))
	}
}

