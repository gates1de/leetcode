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
	k := int(3)
	nums := []int{4,5,8,2}
	addNums := []int{3,5,10,9,4}

	// result: 
	// k := int(0)
	// nums := []int{}
	// addNums := []int{}

	constructor := Constructor(k, nums)
	for _, num := range addNums {
		fmt.Printf("constructor.Add(%v) = %v\n", num, constructor.Add(num))
	}
}
