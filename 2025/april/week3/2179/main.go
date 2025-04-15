package main
import (
	"fmt"
)

type Tree struct {
	Tree []int
}

func constructor(size int) *Tree {
	return &Tree{Tree: make([]int, size + 1)}
}

func (this *Tree) Update(index int, delta int) {
	index++
	for index < len(this.Tree) {
		this.Tree[index] += delta
		index += index & -index
	}
}

func (this *Tree) Query(index int) int {
	index++
	result := int(0)
	for index > 0 {
		result += this.Tree[index]
		index -= index & -index
	}
	return result
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos2 := make([]int, n)
	mapping := make([]int, n)

	for i, num := range nums2 {
		pos2[num] = i
	}

	for i, num := range nums1 {
		mapping[pos2[num]] = i
	}

	tree := constructor(n)
	result := int64(0)

	for value := 0; value < n; value++ {
		pos := mapping[value]
		left := tree.Query(pos)
		tree.Update(pos, 1)
		right := (n - 1 - pos) - (value - left)
		result += int64(left * right)
	}

	return result
}

func main() {
	// result: 1
	// nums1 := []int{2,0,1,3}
	// nums2 := []int{0,1,2,3}

	// result: 4
	nums1 := []int{4,0,1,3,2}
	nums2 := []int{4,1,0,2,3}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := goodTriplets(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

