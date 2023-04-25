package main
import (
	"fmt"
	"sort"
)

type SmallestInfiniteSet struct {
	Current int
	AddedNums []int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{Current: 1, AddedNums: make([]int, 0, 1000)}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	result := this.Current
	if len(this.AddedNums) > 0 {
		result = this.AddedNums[0]
		this.AddedNums = this.AddedNums[1:]
		return result
	}

	this.Current++
	return result
}

func (this *SmallestInfiniteSet) AddBack(num int)  {
	if num >= this.Current {
		return
	}
	this.AddedNums = insert(num, this.AddedNums)
}

func insert(target int, sortedNums []int) []int {
	n := len(sortedNums)
	result := make([]int, n + 1, 1000)

	insertIndex := sort.SearchInts(sortedNums, target)
	if insertIndex >= len(sortedNums) {
		copy(result, sortedNums)
		result[n] = target
		return result
	}

	if target == sortedNums[insertIndex] {
		return sortedNums
	}

	copy(result[:insertIndex], sortedNums[:insertIndex])
	result[insertIndex] = target
	copy(result[insertIndex + 1:], sortedNums[insertIndex:])
	return result
}

func main() {
	// result: [null, null, 1, 2, 3, null, 1, 4, 5]
	controls := [][]int{{1, 2}, {0, 0}, {0, 0}, {0, 0}, {1, 1}, {0, 0}, {0, 0}, {0, 0}}

	// result: 
	// controls := [][]int{{0, 0}, {1, 0}}

	obj := Constructor()
	for _, c := range controls {
		if c[0] == 0 {
			fmt.Printf("obj.PopSmallest() = %v\n", obj.PopSmallest())
		} else if c[0] == 1 {
			obj.AddBack(c[1])
		}
	}
}

